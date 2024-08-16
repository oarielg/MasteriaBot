package discord

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/config"
)

var Session *discordgo.Session

func InitSession() {
	var err error
	// Initializing discord session
	Session, err = discordgo.New("Bot " + config.GetBotToken())
	if err != nil {
		slog.Error("failed to create discord session", "error", err)
	}

	Session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageTyping | discordgo.IntentGuilds
}

func InitConnection() {
	if err := Session.Open(); err != nil { // Creating a connection
		slog.Error("failed to create websocket connection to discord", "error", err)
		return
	}
}
