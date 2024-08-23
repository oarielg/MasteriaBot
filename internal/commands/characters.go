package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
	"github.com/oarielg/MasteriaBot/internal/models"
)

func CreateChar(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) >= 2 {
		if err := models.CreateCharacter(cmd[1], m.Author.ID); err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error creating character!")
		} else {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" character "+cmd[1]+" created successfully!")
		}
	}
}

func ListChars(cmd []string, m *discordgo.MessageCreate) {
	chars, err := models.ListCharacters(m.Author.ID)
	if err != nil {
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error listing characters!")
	} else {
		var fields []*discordgo.MessageEmbedField
		for _, char := range chars {
			fields = append(fields,
				&discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("ID: %d", char.ID),
					Value:  fmt.Sprintf("Character: %s", char.Name),
					Inline: false,
				},
			)
		}

		embed := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Characters",
				Description: "List of your created characters.",
				Fields:      fields,
			},
			},
		}
		discord.SendChannelMessageComplex(m.ChannelID, embed)
	}
}
