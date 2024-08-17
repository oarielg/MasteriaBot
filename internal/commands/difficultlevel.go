package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

func DifficultLevel(cmd []string, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Difficult Level Table",
			Description: "The DL (Difficult Level) defines how hard or easy a Dice Check is.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Routine",
					Value:  "2, 3, 4, 5 and 6",
					Inline: false,
				},
				{
					Name:   "Simple",
					Value:  "3, 4, 5 and 6",
					Inline: false,
				},
				{
					Name:   "Challenging",
					Value:  "4, 5 and 6",
					Inline: false,
				},
				{
					Name:   "Demanding",
					Value:  "5 and 6",
					Inline: false,
				},
				{
					Name:   "Formidable",
					Value:  "6",
					Inline: false,
				},
			},
		},
		},
	}
	discord.SendChannelMessageComplex(m.ChannelID, embed)
}
