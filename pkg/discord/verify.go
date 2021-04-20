package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func (c *client) Verify(m *discordgo.MessageCreate) error {
	userInput := strings.Split(m.Content, " ")

	if len(userInput) < 2 {
		_, err := c.Session.ChannelMessageSend(m.ChannelID, "Please put your profile url after the !verify command")
		if err != nil {
			return err
		}
		return nil
	}

	profileLink := userInput[1]
	splitProfile := strings.Split(profileLink, "/")

	if len(splitProfile) > 3 {
		if splitProfile[3] == "users" || splitProfile[3] == "u" {
			userID, err := strconv.Atoi(splitProfile[4])
			if err != nil {
				return err
			}

			userRank, err := c.Osu.GetUserGlobalRank(userID)
			if err != nil {
				return err
			}

			if userRank < 1000 {
				err := c.Session.GuildMemberRoleAdd(m.Message.GuildID, m.Message.Author.ID, ServerRoles["pro-players"].ID)
				if err != nil {
					c.Session.ChannelMessageSend(m.ChannelID, "You should be able to get this role but adding it failed. Ping an Admin because something awful happened")
					return err
				}
				c.Session.ChannelMessageSend(m.ChannelID, "Thank you. Your Role has been successfully set to Pro-Player!")
				return nil
			} else {
				c.Session.ChannelMessageSend(m.ChannelID, "Sorry! You need to be 3 digit for the Pro-Players role. Please ping an admin for manual inspection")
				return fmt.Errorf("invalid pro player role request")
			}

		} else {
			c.Session.ChannelMessageSend(m.ChannelID, "That is not a valid URL to your profile")
			return fmt.Errorf("invalid profile URL")
		}

	} else {
		c.Session.ChannelMessageSend(m.ChannelID, "That is not a valid URL to your profile")
		return fmt.Errorf("invalid profile URL")
	}
}
