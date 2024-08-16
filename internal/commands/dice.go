package commands

import (
	"errors"
	"fmt"

	"math/rand/v2"
	"regexp"

	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

const pattern = `(\d+)d([\+-]\d+)?$`

func Match(s, p string) ([]string, error) {
	regex, err := regexp.Compile(p)
	if err != nil {
		return nil, err
	}

	if regex.MatchString(s) {
		return regex.FindStringSubmatch(s)[1:], nil
	}

	return nil, errors.New("no matches found for dice roll input: " + s)
}

func RollDice(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		matches, err := Match(cmd[1], pattern)
		if err != nil || len(matches) == 0 {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+cmd[1]+" is not valid!")
			return
		}

		// how many dice?
		dices, err := strconv.Atoi(matches[0])
		if err != nil {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+matches[0]+" is not valid!")
			return
		}

		// modifier?
		var mod int
		if matches[1] != "" {
			mod, err = strconv.Atoi(matches[1])
			if err != nil {
				mod = 0
			}
		}

		var total int
		for i := 0; i < dices; i++ {
			total += Roll(1, 6)
		}

		total = max(total+mod, 0)

		discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s rolled %d!", m.Author.Mention(), total))
	}
}

func RollDamage(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) == 2 {
		if bonus, err := strconv.Atoi(cmd[1]); err == nil {
			damage := max(Roll(1, 6)+bonus, 0)
			discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s rolled %d damage!", m.Author.Mention(), damage))
		} else {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+cmd[1]+" is not a number!")
		}
	}
}

func Roll(min, max int) int {
	return rand.IntN(max-min+1) + min
}
