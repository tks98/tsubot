package discord

import (
	"fmt"
	"github.com/tks98/tsubot/internal/logger"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type client struct {
	Session *discordgo.Session
}

func CreateClient(token string) (*client, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}
	return &client{Session: session}, nil
}

// MessageCreate is called by the AddHandler function everytime a new message is posted in any channel the bot has access too
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore messages sent by the bot
	if m.Author.ID == s.State.User.ID {
		return
	}

	logger.Log.Debugf("Got message %s", m.Content)

	// Check if the message sent is a valid key in the commands map
	// If so, calls the function for that command
	if command, ok := Commands[strings.ToLower(m.Content)]; ok {
		err := command(m)
		if err != nil {
			logger.Log.Error(err)
		}
	}
}
