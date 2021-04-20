package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
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
		if role.Name == "Pro-Players" {
			if _, err := c.Session.ChannelMessageSend(m.ChannelID, "You need to verify your osu! account. Please type !verify followed by a link to your account"); err != nil {
				return err
			}
		} else {
			if message[0] == "!choose" {
				err := c.Session.GuildMemberRoleAdd(m.Message.GuildID, m.Message.Author.ID, role.ID)
				if err != nil {
					return err
				}

				content := fmt.Sprintf("Your role was successfuly changed to %s", role.Name)
				if _, err := c.Session.ChannelMessageSend(m.ChannelID, content); err != nil {
					return err
				}
			} else if message[0] == "!remove" {
				err := c.Session.GuildMemberRoleRemove(m.Message.GuildID, m.Message.Author.ID, role.ID)
				if err != nil {
					return err
				}

				content := fmt.Sprintf("You no longer have the %s role", role.Name)
				if _, err := c.Session.ChannelMessageSend(m.ChannelID, content); err != nil {
					return err
				}
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

func (c *client) ListRoles(m *discordgo.MessageCreate) error {

	var roles []string
	for role, _ := range AllowedRoles {
		roles = append(roles, role)
	}
	content := fmt.Sprintf("These are the roles you can choose with !roles:  %v", roles)

	if _, err := c.Session.ChannelMessageSend(m.ChannelID, content); err != nil {
		return err
	}

	return nil
}
