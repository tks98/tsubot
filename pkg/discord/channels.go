package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var GeneralID string

func (c *client) GetGeneralChannelID() (string, error) {

	channels, err := c.Session.GuildChannels(c.GuildID)
	if err != nil {
		return "", err
	}

	for _, channel := range channels {
		if channel.Name == "general" && channel.Type == discordgo.ChannelTypeGuildText {
			GeneralID = channel.ID
			return channel.ID, nil
		}
	}

	return "", fmt.Errorf("could not find general channel id")
}
