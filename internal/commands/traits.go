package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
	"github.com/oarielg/MasteriaBot/internal/models"
)

func ListTraits(cmd []string, m *discordgo.MessageCreate) {
	traits, err := models.ListTraits()
	if err != nil {
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error listing Traits!")
	} else {
		var body string
		body += "```\n"
		for _, trait := range traits {
			body += fmt.Sprintf("%s, ID: %d\n", trait.Name, trait.ID)
		}
		body += "```"
		discord.SendChannelMessage(m.ChannelID, body)
	}
}

func ShowTrait(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		trait, err := models.GetTrait(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" Trait not found!")
		} else {
			var body string
			body += "```\n"
			body += fmt.Sprintf("%s\n%s", trait.Name, trait.Description)
			body += "```"
			discord.SendChannelMessage(m.ChannelID, body)
		}
	}
}
