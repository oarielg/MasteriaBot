package discord

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

func SearchGuildByChannelID(textChannelID string) (guildID string) {
	channel, _ := Session.Channel(textChannelID)
	guildID = channel.GuildID
	return guildID
}

func SendChannelMessage(channelID string, message string) {
	_, err := Session.ChannelMessageSend(channelID, message)
	if err != nil {
		slog.Warn("failed to send message to channel", "channelId", channelID, "message", message, "error", err)
	}
}

func SendChannelMessageComplex(channelID string, message *discordgo.MessageSend) {
	_, err := Session.ChannelMessageSendComplex(channelID, message)
	if err != nil {
		slog.Warn("failed to send message to channel", "channelId", channelID, "message", message, "error", err)
	}
}
