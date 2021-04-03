package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/internal/logger"
	"time"
)

var Commands map[string]func(*discordgo.MessageCreate) error

func (c *client) InitCommands(commands []string) error {

	Commands = make(map[string]func(*discordgo.MessageCreate) error)

	for _, command := range commands {
		if command == "!ping" {
			Commands[command] = c.ping
		}
	}

	logger.Log.Info(Commands)

	return nil
}

func (c *client) ping(m *discordgo.MessageCreate) error {
	start := time.Now()
	c.Session.ChannelMessageSend(m.ChannelID, "I'm working, stop poking me")
	stop := time.Now()
	c.Session.ChannelMessageSend(m.ChannelID, "Latency: "+stop.Sub(start).String())

	return nil
}
