package handlers

import (
	"log/slog"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/commands"
	"github.com/oarielg/MasteriaBot/internal/config"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

const help string = "This is help"

func ReadyHandler(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status.
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

	//	Splitting command into string slice
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
	default:
		return
	}
}
