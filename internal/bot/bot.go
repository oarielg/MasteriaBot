package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oarielg/MasteriaBot/internal/config"
	"github.com/oarielg/MasteriaBot/internal/database"
	"github.com/oarielg/MasteriaBot/internal/discord"
	"github.com/oarielg/MasteriaBot/internal/handlers"
)

func Run() {
	config.Load()

	//create a session
	discord.InitSession()

	//add handlers
	addHandlers()

	//open session
	discord.InitConnection()
	defer discord.Session.Close()

	// init database
	database.Connect()

	fmt.Println("Bot running....")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func addHandlers() {
	discord.Session.AddHandler(handlers.ReadyHandler)
	discord.Session.AddHandler(handlers.MessageCreateHandler)
}
