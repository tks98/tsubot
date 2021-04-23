package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	oppai "github.com/flesnuk/oppai5"
	"github.com/tks98/tsubot/internal/util"
	"github.com/tks98/tsubot/pkg/osu"
	"os"
	"strings"
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
	mods := score.Mods
	if len(mods) == 0 {
		mods = append(mods, "NM")
	}

	file, err := c.Osu.DownloadBeatmapFile(fmt.Sprintf("%d", score.Beatmap.ID))
	if err != nil {
		return nil, err
	}

	defer os.Remove(file.Name())

	// convert the mods array [HD,HR,DT] into the bitwise mod combination for the performance calculator
	modsBit := uint32(osu.ParseMods(strings.Join(score.Mods[:], "")))
	parameters := &oppai.Parameters{
		Mods:   modsBit,
		Combo:  uint16(score.MaxCombo),
		N300:   uint16(score.Statistics.Count300),
		N100:   uint16(score.Statistics.Count100),
		N50:    uint16(score.Statistics.Count50),
		Misses: uint16(score.Statistics.CountMiss),
	}

	// calculate the performance point information for the score
	performance, err := c.Osu.PerformanceCalc(file, parameters)
	if err != nil {
		return nil, err
	}

	// create the embed to display score information
	embed := &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s [%s]", score.Beatmapset.Title, score.Beatmap.Version),
		URL:   score.Beatmap.URL,

		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: score.Beatmapset.Covers.List2X,

		},

		Fields: []*discordgo.MessageEmbedField{
			{
				Inline: true,
				Name:   "PP",
				Value:  fmt.Sprintf("%.2f", score.Pp),
			},

			{
				Inline: true,
				Name:   "Stars",
				Value:  fmt.Sprintf("%.2f", performance.Pp.Diff.Total),
			},

			{
				Inline: true,
				Name:   "Length",
				Value:  util.SecondsToMinutes(score.Beatmap.TotalLength),
			},

			{
				Inline: true,
				Name:   "BPM",
				Value:  fmt.Sprintf("%.2f", score.Beatmap.Bpm),
			},
			{
				Inline: true,
				Name:   "Mods",
				Value:  fmt.Sprintf("%v", mods),
			},
			{
				Inline: true,
				Name:   "Accuracy",
				Value:  fmt.Sprintf("%.2f%s", score.Accuracy * 100, "%"),
			},
			{
				Inline: true,
				Name:   "Acc Stats",
				Value:  fmt.Sprintf("[%d/%d/%d/%d]", score.Statistics.Count300, score.Statistics.Count100, score.Statistics.Count50, score.Statistics.CountMiss),
			},
			{
				Inline: true,
				Name:   "Combo",
				Value:  fmt.Sprintf("%d/%dx", score.MaxCombo, performance.BeatmapInfo.MaxCombo),
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
