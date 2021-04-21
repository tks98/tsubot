package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (c *client) GetOsuStat(m *discordgo.MessageCreate) error {
	content := strings.Split(m.Content, " ")
	stat := content[0]

	switch stat {
	case "!rank":
		if err := c.getRank(m); err != nil {
			return err
		}
	}

	return nil
}

func (c *client) getRank(m *discordgo.MessageCreate) error {
	content := strings.Split(m.Content, " ")
	rank, err := c.Osu.GetUserGlobalRank(content[1])
	if err != nil {
		return err
	}
	reply := fmt.Sprintf("Rank: %d", rank)
	if _, err := c.Session.ChannelMessageSend(m.ChannelID, reply); err != nil {
		return err
	}

	return nil
}


