package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/internal/logger"
)

func MemberJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {

	err := s.GuildMemberRoleAdd(member.GuildID, member.User.ID, ServerRoles["gamer"].ID)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	_, err = s.ChannelMessageSend(GeneralID, "Welcome to LRS! "+member.Mention()+" Your role has been set to osu! player. Please use the !roles command in the #roles channel for a list you can choose from.")
	if err != nil {
		logger.Log.Error(err)
	}
}
