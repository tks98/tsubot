package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/pkg/osu"
)

type client struct {
	Session *discordgo.Session
	GuildID string
	Osu     *osu.Client
}

func CreateClient(token, guildID, clientID, clientSecret string) (*client, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, err
	}

	osuClient, err := osu.NewClient(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	return &client{Session: session, GuildID: guildID, Osu: osuClient}, nil
}
