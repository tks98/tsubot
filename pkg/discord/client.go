package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type client struct {
	Session  *discordgo.Session
	Commands map[string]func(discordgo.Session) error
}

func CreateClient(token string) (*client, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}

	return &client{Session: session}, nil
}

// MessageCreate is called by the AddHandler function everytime a new message is posted in any channel the bot has access too
func (c *client) MessageCreate(m *discordgo.MessageCreate) error {

	// Ignore messages sent by the bot
	if m.Author.ID == c.Session.State.User.ID {
		return nil
	}

	// Check if the message send is a valid key in the commands map
	// If so, calls the function for that command
	if command, ok := c.Commands[strings.ToLower(m.Content)]; ok {
		err := command(*c.Session)
		if err != nil {
			return err
		}
	}

	return nil

}
