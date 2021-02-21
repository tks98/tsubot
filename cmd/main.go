package main

import (
	"github.com/tks98/tsubot/config"
	"github.com/tks98/tsubot/internal/logger"
	"github.com/tks98/tsubot/pkg/discord"
)

func main() {

	// Parse config and init logger
	config := config.GetConfigs()
	logger.InitLogger(nil)
	logger.Log.Info("Config parsed and logger init")

	// Create discord client
	client, err := discord.CreateClient(config.APIKeys.Discord)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Register handlers
	client.Session.AddHandler(client.MessageCreate)

}
