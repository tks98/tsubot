package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/internal/logger"
	"strings"
)

var ServerRoles map[string]*discordgo.Role
var AllowedRoles map[string]struct{}

func (c *client) InitRoles() error {

	// Get all of the roles on the server
	roles, err := c.Session.GuildRoles(c.GuildID)
	if err != nil {
		return err
	}

	// Only add the allowed roles to the ServerRoles Map
	ServerRoles = make(map[string]*discordgo.Role)
	for _, role := range roles {
		if _, ok := AllowedRoles[strings.ToLower(role.Name)]; ok {
			ServerRoles[strings.ToLower(role.Name)] = role
		}
	}

	return nil
}

func (c *client) ChangeRole(m *discordgo.MessageCreate) error {
	message := strings.Split(strings.ToLower(m.Content), " ")

	if role, ok := ServerRoles[message[1]]; ok {
		logger.Log.Info("here")
		logger.Log.Info(role.Name)
		if role.Name == "Pro-Players" {
			if _, err := c.Session.ChannelMessageSend(m.ChannelID, "You need to verify your osu! account. Please type !verify followed by a link to your account"); err != nil {
				return err
			}
		} else {
			err := c.Session.GuildMemberRoleAdd(m.Message.GuildID, m.Message.Author.ID, role.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *client) SetAllowedRoles(roles []string) {
	AllowedRoles = make(map[string]struct{})
	for _, role := range roles {
		AllowedRoles[strings.ToLower(role)] = struct{}{}
	}
}
