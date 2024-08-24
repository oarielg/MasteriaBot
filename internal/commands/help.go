package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

func Help(cmd []string, m *discordgo.MessageCreate) {
	var body string

	body += "```\n"
	body += "Commands:\n"
	body += "!dl\n"
	body += "!roll <dice> (Example: !roll 1d)\n"
	body += "!dam <bonus> (Example: !dam 2)\n"
	body += "!dc <difficulty> <modifier> (Example: !dc challenging)\n"
	body += "!createchar <char_name> (Use underscores instead of spaces. Example: !createchar Dark_Schneider)\n"
	body += "!deletechar <char_id>\n"
	body += "!chars\n"
	body += "!char <char_id>\n"
	body += "!changename <char_id> <new_name> (Use underscores instead of spaces. Example: !changename 3 Dark_Schneider)\n"
	body += "!traits\n"
	body += "!trait <trait_id>\n"
	body += "!powers\n"
	body += "!power <power_id>\n"
	body += "!addtrait <char_id> <trait_id>\n"
	body += "!removetrait <char_id> <trait_id>\n"
	body += "!addpower <char_id> <power_id>\n"
	body += "!removepower <char_id> <power_id>\n"
	body += "```"

	discord.SendChannelMessage(m.ChannelID, body)
}
