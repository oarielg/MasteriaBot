package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/oarielg/MasteriaBot/internal/discord"
)

type difficultLevel int

const (
	_ difficultLevel = iota
	None
	Routine
	Simple
	Challenging
	Demanding
	Formidable
)

var difficultLevelToString = map[difficultLevel]string{
	None:        "None",
	Routine:     "Routine",
	Simple:      "Simple",
	Challenging: "Challenging",
	Demanding:   "Demanding",
	Formidable:  "Formidable",
}

var stringToDifficultLevel = map[string]difficultLevel{
	"routine":     Routine,
	"simple":      Simple,
	"challenging": Challenging,
	"demanding":   Demanding,
	"formidable":  Formidable,
}

func toDifficultLevel(s string) difficultLevel {
	if dl, ok := stringToDifficultLevel[s]; ok {
		return dl
	}
	return None
}

func (dl difficultLevel) String() string {
	if s, ok := difficultLevelToString[dl]; ok {
		return s
	}
	return "unknown"
}

func DiceCheck(cmd []string, m *discordgo.MessageCreate) {
	if len(cmd) >= 2 {
		dl := toDifficultLevel(strings.ToLower(cmd[1]))
		if dl == None {
			discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+cmd[1]+" is not a valid Difficult Level!")
			return
		}

		dl_s := dl.String()
		var mod string
		if len(cmd) == 3 {
			cmd[2] = strings.ToLower(cmd[2])
			switch cmd[2] {
			case "s":
				dl = dl - 2
				mod = "superiority"
			case "a":
				dl = dl - 1
				mod = "advantage"
			case "d":
				dl = dl + 1
				mod = "disadvantage"
			case "i":
				dl = dl + 2
				mod = "inferiority"
			default:
				discord.SendChannelMessage(m.ChannelID, m.Author.Mention()+" "+cmd[2]+" is not a valid Modifier!")
				return
			}
			if dl <= 1 { // automatic success
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check with %s, but it's an automatic success!", m.Author.Mention(), dl_s, mod))
				return
			}
			if dl > 6 { // automatic failure
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check with %s, but it's an automatic failure!", m.Author.Mention(), dl_s, mod))
				return
			}
		}

		roll := Roll(1, 6)
		if roll >= int(dl) {
			if mod != "" {
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check with %s and rolled %d; a success!", m.Author.Mention(), dl_s, mod, roll))
			} else {
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check and rolled %d; a success!", m.Author.Mention(), dl_s, roll))
			}
		} else {
			if mod != "" {
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check with %s and rolled %d; a failure!", m.Author.Mention(), dl_s, mod, roll))
			} else {
				discord.SendChannelMessage(m.ChannelID, fmt.Sprintf("%s is doing a DL %s dice check and rolled %d; a failure!", m.Author.Mention(), dl_s, roll))
			}
		}
	}
}
