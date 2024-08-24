package handlers

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/commands"
	"github.com/oarielg/MasteriaBot/internal/config"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status
	err := s.UpdateGameStatus(0, config.GetBotStatus())
	if err != nil {
		slog.Warn("failed to update game status", "error", err)
	}
}

func MessageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Preventing bot from replying to itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	slog.Info("processing command", "command", m.Content)

	prefix := config.GetBotPrefix()
	//guildID := discord.SearchGuildByChannelID(m.ChannelID)

	cmd := strings.Split(m.Content, " ")

	switch cmd[0] {
	case prefix + "help":
		commands.Help(cmd, m)
	case prefix + "ping":
		discord.SendChannelMessage(m.ChannelID, "Pong!")
	case prefix + "roll":
		commands.RollDice(cmd, m)
	case prefix + "dam":
		commands.RollDamage(cmd, m)
	case prefix + "dc":
		commands.DiceCheck(cmd, m)
	case prefix + "dl":
		commands.DifficultLevel(cmd, m)
	case prefix + "createchar":
		commands.CreateChar(cmd, m)
	case prefix + "deletechar":
		commands.DeleteChar(cmd, m)
	case prefix + "chars":
		commands.ListChars(cmd, m)
	case prefix + "char":
		commands.ShowChar(cmd, m)
	case prefix + "changename":
		commands.ChangeName(cmd, m)
	case prefix + "addtrait":
		commands.AddTrait(cmd, m)
	case prefix + "removetrait":
		commands.RemoveTrait(cmd, m)
	case prefix + "addpower":
		commands.AddPower(cmd, m)
	case prefix + "removepower":
		commands.RemovePower(cmd, m)
	case prefix + "traits":
		commands.ListTraits(cmd, m)
	case prefix + "trait":
		commands.ShowTrait(cmd, m)
	case prefix + "powers":
		commands.ListPowers(cmd, m)
	case prefix + "power":
		commands.ShowPower(cmd, m)
	default:
		return
	}
}
