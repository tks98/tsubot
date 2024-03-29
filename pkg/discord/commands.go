package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/internal/logger"
	"strings"
)

var Commands map[string]interface{}

const prefix = "!"

func (c *client) InitCommands(commands []string) error {

	Commands = make(map[string]interface{})

	for _, command := range commands {
		switch command {
		case "ping":
			Commands[prefix+command] = c.Ping
		case "choose":
			Commands[prefix+command] = c.ChangeRole
		case "verify":
			Commands[prefix+command] = c.Verify
		case "roles":
			Commands[prefix+command] = c.ListRoles
		case "remove":
			Commands[prefix+command] = c.ChangeRole
		case "info":
			Commands[prefix+command] = c.GetOsuStat
		}
	}

	logger.Log.Info(Commands)

	return nil
}

// HandleMessage is called by the AddHandler function everytime a new message is posted in any channel the bot has access too
func HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages sent by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	//logger.Log.Debugf("Got message %s", m.Content)
	cmd := strings.Split(strings.ToLower(m.Content), " ")[0]
	//logger.Log.Debugf(cmd)

	// Check if the first string in the message is a valid key in the commands map
	// If so, calls the function for that command
	if command, ok := Commands[cmd]; ok {
		logger.Log.Infof("Valid Command %s", m.Content)
		err := command.(func(*discordgo.MessageCreate) error)(m)
		if err != nil {
			logger.Log.Error(err)
		}
	}
}

func (c *client) Ping(m *discordgo.MessageCreate) error {
	content := fmt.Sprintf("I'm working, stop poking me: %s", c.Session.HeartbeatLatency())
	if _, err := c.Session.ChannelMessageSend(m.ChannelID, content); err != nil {
		return err
	}

	return nil
}
