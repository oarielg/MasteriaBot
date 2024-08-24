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

func AddTrait(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 3 {
		char, err := models.GetCharacter(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this character does not exist!")
			return
		}
		if char.Owner != m.Author.ID {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this action is not allowed!")
			return
		}

		trait, err := models.GetTrait(cmd[2])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" Trait not found!")
			return
		}

		err = models.AddCharacterTrait(char, trait)
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" an error ocurred while adding the Trait to the character!")
		} else {
			discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s, %s trait added to %s successfully!", m.Author.Mention(), trait.Name, char.Name))
		}
	}
}

func RemoveTrait(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 3 {
		char, err := models.GetCharacter(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this character does not exist!")
			return
		}
		if char.Owner != m.Author.ID {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this action is not allowed!")
			return
		}

		trait, err := models.GetTrait(cmd[2])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" Trait not found!")
			return
		}

		err = models.RemoveCharacterTrait(char, trait)
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" an error ocurred while removing the Trait from the character!")
		} else {
			discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s, %s trait removed from %s successfully!", m.Author.Mention(), trait.Name, char.Name))
		}
	}
}
