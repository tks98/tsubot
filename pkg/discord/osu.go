package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (c *client) GetOsuStat(m *discordgo.MessageCreate) error {
	content := strings.Split(strings.ToLower(m.Content), " ")
	var userID, stat, offset string
	if len(content) == 3 {
		userID = content[2] // woey
		stat = content[1]   // -top, recent, firsts

		if stat == "-recent" || stat == "-r" {
			offset = "1"
			stat = "recent"
		} else if stat == "-top" || stat == "-t" {
			offset = "100"
			stat = "best"
		} else if stat == "-firsts" || stat == "-f" {
			offset = "100"
			stat = "firsts"
		} else {
			return fmt.Errorf("you did not enter valid flags for the !info command")
		}

		scores, err := c.Osu.GetUserScores(userID, stat, offset)
		if err != nil {
			return err
		}

		embed, err := c.createRecentScoreEmbed(scores)
		if err != nil {
			return err
		}

		err = c.PostEmbed(m.ChannelID, embed)
		if err != nil {
			return err
		}


	} else if len(content) == 2 {
		userID = content[1]
		user, err := c.Osu.GetUser(userID)
		if err != nil {

		}
		err = c.PostEmbed(m.ChannelID, c.createUserInfoEmbed(user))
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("you did not enter valid flags for the !info command")
	}

	return nil
}

func (c *client) getRank(m *discordgo.MessageCreate) error {
	content := strings.Split(m.Content, " ")
	rank, err := c.Osu.GetUser(content[1])
	if err != nil {
		return err
	}
	reply := fmt.Sprintf("Rank: %d", rank)
	if _, err := c.Session.ChannelMessageSend(m.ChannelID, reply); err != nil {
		return err
	}

	return nil
}
