package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
	"github.com/oarielg/MasteriaBot/internal/models"
)

func ListPowers(cmd []string, m *discordgo.MessageCreate) {
	powers, err := models.ListPowers()
	if err != nil {
		discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" error listing Powers!")
	} else {
		var body string
		body += "```\n"
		for _, power := range powers {
			body += fmt.Sprintf("%s, ID: %d\n", power.Name, power.ID)
		}
		body += "```"
		discord.SendChannelMessage(m.ChannelID, body)
	}
}

func ShowPower(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		power, err := models.GetPower(cmd[1])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" Power not found!")
		} else {
			var body string
			body += "```\n"
			body += fmt.Sprintf("%s\n%s Willpower, %s\n%s", power.Name, power.Willpower, power.Duration, power.Description)
			body += "```"
			discord.SendChannelMessage(m.ChannelID, body)
		}
	}
}
