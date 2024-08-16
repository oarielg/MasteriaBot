package discord

import (
	"log/slog"
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
