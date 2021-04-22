package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/tks98/tsubot/internal/util"
	"github.com/tks98/tsubot/pkg/osu"
)

func (c *client) createUserInfoEmbed(user *osu.User) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Title: user.Username,
		URL:   fmt.Sprintf("https://osu.ppy.sh/users/%d/", user.ID),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:    user.AvatarURL,
			Width:  500,
			Height: 500,
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "tsubot",
			IconURL: "https://cdn.discordapp.com/attachments/611191473601511434/834593625514049536/botimage.jpg",
		},

		Fields: []*discordgo.MessageEmbedField{
			{
				Inline: true,
				Name:   "Global Rank",
				Value:  util.NumberToString(user.Statistics.GlobalRank, ','),
			},

			{
				Inline: true,
				Name:   "PP",
				Value:  util.NumberToString(int(user.Statistics.Pp), ','),
			},
			{
				Inline: true,
				Name:   "Ranked Score",
				Value:  util.NumberToString(int(user.Statistics.RankedScore), ','),
			},
			{
				Inline: true,
				Name:   "Total Hits",
				Value:  util.NumberToString(int(user.Statistics.TotalHits), ','),
			},
			{
				Inline: true,
				Name:   "Max Combo",
				Value:  util.NumberToString(int(user.Statistics.MaximumCombo), ','),
			},
			{
				Inline: true,
				Name:   "Play Count",
				Value:  util.NumberToString(int(user.Statistics.PlayCount), ','),
			},
		},
	}

	return embed

}

func (c *client) createRecentScoreEmbed(scores *osu.UserScores) (*discordgo.MessageEmbed, error) {

	if len(*scores) == 0 {
		return nil, fmt.Errorf("scores api call returned empty")

	}

	score := (*scores)[1]

	embed := &discordgo.MessageEmbed{
		Title: score.Beatmapset.Title,
		URL:   score.Beatmap.URL,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: score.Beatmapset.Covers.Slimcover2X,
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "tsubot",
			IconURL: "https://cdn.discordapp.com/attachments/611191473601511434/834593625514049536/botimage.jpg",
		},

		Fields: []*discordgo.MessageEmbedField{
			{
				Inline: true,
				Name:   "PP",
				Value:  fmt.Sprintf("%f", score.Pp),
			},
			{
				Inline: true,
				Name:   "Length",
				Value:  util.NumberToString(score.Beatmap.TotalLength, ','),
			},

			{
				Inline: true,
				Name:   "BPM",
				Value:  util.NumberToString(score.Beatmap.Bpm, ','),
			},
			{
				Inline: true,
				Name:   "Mods",
				Value:  fmt.Sprintf("%v", score.Mods),
			},
			{
				Inline: true,
				Name:   "Accuracy",
				Value:  fmt.Sprintf("%f", score.Accuracy),
			},
			{
				Inline: true,
				Name:   "Acc Stats",
				Value:  fmt.Sprintf("[%d/%d/%d/%d]", score.Statistics.Count300, score.Statistics.Count100, score.Statistics.Count50, score.Statistics.CountMiss),
			},
			{
				Inline: true,
				Name:   "Max Combo",
				Value:  fmt.Sprintf("%d", score.MaxCombo),
			},
		},
	}

	return embed, nil

}

func (c *client) PostEmbed(channelID string, embed *discordgo.MessageEmbed) error {

	_, err := c.Session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: embed,
	})

	if err != nil {
		return err
	}
	return nil
}
