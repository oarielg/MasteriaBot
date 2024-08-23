package handlers

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/commands"
	"github.com/oarielg/MasteriaBot/internal/config"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

const help string = "**Commands:**\n!ping\n!dl\n!roll <dice> *(Example: !roll 1d)*\n!dam <bonus> *(Example: !dam 2)*\n!dc <difficulty> <modifier> *(Example: !dc challenging)*"

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
		discord.SendChannelMessage(m.ChannelID, help)
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
	case prefix + "chars":
		commands.ListChars(cmd, m)
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
