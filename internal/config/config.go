package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type configuration struct {
	BotPrefix string
	BotStatus string
	BotToken  string
}

var config *configuration

func Load() {
	err := godotenv.Load("local.env")
	if err != nil {
		slog.Error("failed to load environment variables")
	}

	config = &configuration{
		BotPrefix: os.Getenv("BOT_PREFIX"),
		BotStatus: os.Getenv("BOT_STATUS"),
		BotToken:  os.Getenv("BOT_TOKEN"),
	}
}

func GetBotPrefix() string {
	return config.BotPrefix
}

func GetBotStatus() string {
	return config.BotStatus
}

func GetBotToken() string {
	return config.BotToken
}
