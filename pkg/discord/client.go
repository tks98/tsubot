package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type client struct {
	Session *discordgo.Session
	GuildID string
}

func CreateClient(token string, guildID string) (*client, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}
	return &client{Session: session, GuildID: guildID}, nil
}
