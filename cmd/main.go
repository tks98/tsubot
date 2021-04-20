package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/config"
	"github.com/tks98/tsubot/internal/logger"
	"github.com/tks98/tsubot/pkg/discord"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Parse config and init logger
	config := config.GetConfigs()
	logger.InitLogger(nil)
	logger.Log.Info("Config parsed and logger init")

	// Create discord client
	client, err := discord.CreateClient(config.APIKeys.Discord, config.GuideID, config.APIKeys.Osu.ClientID, config.APIKeys.Osu.ClientSecret)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Retrieve the ID of the general channel
	_, err = client.GetGeneralChannelID()
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Parse config and init supported commands
	err = client.InitCommands(config.Commands)
	if err != nil {
		logger.Log.Fatal(err)
	}

	client.SetAllowedRoles(config.AllowedRoles)

	// Parse config and init supported commands
	err = client.InitRoles()
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Register handlers
	client.Session.AddHandler(discord.HandleMessage)
	client.Session.AddHandler(discord.MemberJoin)

	client.Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// Open a websocket connection to Discord and begin listening.
	err = client.Session.Open()
	if err != nil {
		logger.Log.Fatal(err)
		return
	}

	// Bot runs until term signal is received
	logger.Log.Info("tsubot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = client.Session.Close()
	if err != nil {
		logger.Log.Fatal(err)
	}

}
