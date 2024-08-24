package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/database"
	"github.com/oarielg/MasteriaBot/internal/discord"
	"github.com/oarielg/MasteriaBot/internal/models"
)

func CreateChar(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) >= 2 {
		name := strings.ReplaceAll(cmd[1], "_", " ")
		if name == "" {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" the name can't be empty!")
			return
		}

		if err := models.CreateCharacter(name, m.Author.ID); err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error creating character!")
		} else {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" character "+name+" created successfully!")
		}
	}
}

func ListChars(cmd []string, m *discordgo.MessageCreate) {
	chars, err := models.ListCharacters(m.Author.ID)
	if err != nil {
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error listing characters!")
	} else {
		var body string
		body += m.Author.Mention() + ", **your Characters:**\n"
		body += "```\n"

		for _, char := range chars {
			body += fmt.Sprintf("-%s, ID: %d\n", char.Name, char.ID)
		}

		body += "```"
		discord.SendChannelMessage(m.ChannelID, body)
	}
}

func ShowChar(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		char, err := models.GetCharacter(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this character does not exist!")
		} else {
			var body, traits, powers string
			for _, trait := range char.Traits {
				traits += fmt.Sprintf("%s (ID: %d), ", trait.Name, trait.ID)
			}
			traits = strings.TrimSuffix(traits, ", ")
			for _, power := range char.Powers {
				powers += fmt.Sprintf("%s (ID: %d), ", power.Name, power.ID)
			}
			powers = strings.TrimSuffix(powers, ", ")

			body += "```\n"
			body += fmt.Sprintf("Name: %s\nVitality: %d\nWillpower: %d\nTraits: %s\nPowers: %s", char.Name, char.Vitality, char.Willpower, traits, powers)
			body += "```"
			discord.SendChannelMessage(m.ChannelID, body)
		}
	}
}

func ChangeName(cmd []string, m *discordgo.MessageCreate) {
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

		name := strings.ReplaceAll(cmd[2], "_", " ")
		if name == "" {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" the name can't be empty!")
			return
		}

		oldname := char.Name
		char.Name = name
		database.DB.Save(&char)
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+", "+oldname+" has been renamed to "+name+".")
	}
}

func DeleteChar(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		char, err := models.GetCharacter(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this character does not exist!")
			return
		}
		if char.Owner != m.Author.ID {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" this action is not allowed!")
			return
		}

		database.DB.Delete(&char)
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+", character "+char.Name+" has been successfully deleted.")
	}
}
