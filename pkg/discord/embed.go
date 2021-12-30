package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	oppai "github.com/flesnuk/oppai5"
	"github.com/tks98/tsubot/internal/util"
	"github.com/tks98/tsubot/pkg/osu"
	"os"
	"strings"
	"time"
)

func (c *client) createUserInfoEmbed(user *osu.User) *discordgo.MessageEmbed {

	title := fmt.Sprintf("osu! Standard Profile Stats for %s", user.Username)

	description := fmt.Sprintf("> **Rank:** %s\n > **Level:** %d => %d%% to %d\n > **PP:** %.2f **Acc**: %.2f%%\n > **Playcount:** %d (%d hours)\n > **Ranked Score:** %s",
		util.NumberToString(user.Statistics.GlobalRank, ','), user.Statistics.Level.Current, user.Statistics.Level.Progress, user.Statistics.Level.Current+1, user.Statistics.Pp, user.Statistics.HitAccuracy, user.Statistics.PlayCount, user.Statistics.PlayTime/3600, util.NumberToString(int(user.Statistics.RankedScore), ','))

	var status string
	if user.IsOnline {
		status = fmt.Sprintf("Currently Online")
	} else {
		status = fmt.Sprintf("User is currently offline")
	}

	embed := &discordgo.MessageEmbed{

		Author: &discordgo.MessageEmbedAuthor{
			Name:    title,
			URL:     fmt.Sprintf("https://osu.ppy.sh/users/%d/", user.ID),
			IconURL: "https://upload.wikimedia.org/wikipedia/commons/4/44/Osu%21Logo_%282019%29.png",
		},

		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: user.AvatarURL,
		},

		Description: description,

		Footer: &discordgo.MessageEmbedFooter{
			Text: status,
		},
	}

	return embed

}

func (c *client) createRecentScoreEmbed(scores *osu.UserScores, offset string) (*discordgo.MessageEmbed, error) {

	if len(*scores) == 0 {
		return nil, fmt.Errorf("scores api call returned empty")

	}

	score := (*scores)[0]
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
	title := fmt.Sprintf("%s [%s] +%s [%s, %.2f★] | (%.2f%%) %d/%dx | %.2fpp",
		score.Beatmapset.Title, score.Beatmap.Version, mods, score.Beatmapset.Creator,
		performance.Pp.Diff.Total, score.Accuracy*100, score.MaxCombo,
		performance.BeatmapInfo.MaxCombo, score.Pp)

	mapInfo := fmt.Sprintf("AR: **%.1f** | OD: **%.1f** | CS: **%.1f** | HP: **%.1f** | Length: **%s**", performance.Pp.Stats.AR, performance.Pp.Stats.OD, performance.Pp.Stats.CS, performance.Pp.Stats.HP, util.SecondsToMinutes(score.Beatmap.TotalLength))

	accStats := fmt.Sprintf("[**%d**/**%d**/**%d**/**%d**]", score.Statistics.Count300, score.Statistics.Count100, score.Statistics.Count50, score.Statistics.CountMiss)

	// try to determine if the score was an fc
	var ifFC string
	var description string
	if performance.BeatmapInfo.MaxCombo-score.MaxCombo < 10 {
		ifFC = fmt.Sprintf("**%.2fpp** for __%.2f%%__", performance.PpFc.PP.Total, performance.PpFc.PP.ComputedAccuracy.Value()*100)
		description = fmt.Sprintf("> **Map:** %s\n > **Acc:** %s\n > **FC:** %s", mapInfo, accStats, ifFC)
	} else if score.Perfect {
		description = fmt.Sprintf("> **Map:** %s\n > **Acc:** %s\n", mapInfo, accStats)
	} else {
		ifFC = fmt.Sprintf("**%.2fpp** for __%.2f%%__", performance.PpFc.PP.Total, performance.PpFc.PP.ComputedAccuracy.Value()*100)
		description = fmt.Sprintf("> **Map:** %s\n > **Acc:** %s\n > **FC:** %s", mapInfo, accStats, ifFC)
	}

	// calculate the time since the score was set
	duration := time.Since(score.CreatedAt)

	var since float64
	var durationScale string
	var footer string
	if duration.Seconds() > 60 && duration.Seconds() < 3600 {
		since = duration.Minutes()
		durationScale = "minutes"
		footer = fmt.Sprintf("Set %.0f %s ago", since, durationScale)
	} else if duration.Seconds() > 3600 && duration.Seconds() < 86400 {
		since = duration.Hours()
		durationScale = "hours"
		footer = fmt.Sprintf("Set %.1f %s ago", since, durationScale)
	} else if duration.Seconds() > 86400 && duration.Seconds() < 31556952 {
		since = duration.Hours() / 24
		durationScale = "days"
		footer = fmt.Sprintf("Set %.0f %s ago", since, durationScale)
	} else if duration.Seconds() > 31556952 {
		since = duration.Hours() / 8760
		durationScale = "years"
		footer = fmt.Sprintf("Set %.1f %s ago", since, durationScale)
	} else {
		since = duration.Seconds()
		durationScale = "seconds"
		footer = fmt.Sprintf("Set %.0f %s ago", since, durationScale)
	}

	// create the embed to display score information
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    title,
			URL:     score.Beatmap.URL,
			IconURL: score.User.AvatarURL,
		},

		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: score.Beatmapset.Covers.List2X,
		},

		Description: description,

		Footer: &discordgo.MessageEmbedFooter{
			Text:    footer,
			IconURL: "https://upload.wikimedia.org/wikipedia/commons/4/44/Osu%21Logo_%282019%29.png",
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
