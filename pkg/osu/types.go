package osu

import "time"

type User struct {
	AvatarURL     string      `json:"avatar_url"`
	CountryCode   string      `json:"country_code"`
	DefaultGroup  string      `json:"default_group"`
	ID            int         `json:"id"`
	IsActive      bool        `json:"is_active"`
	IsBot         bool        `json:"is_bot"`
	IsDeleted     bool        `json:"is_deleted"`
	IsOnline      bool        `json:"is_online"`
	IsSupporter   bool        `json:"is_supporter"`
	LastVisit     interface{} `json:"last_visit"`
	PmFriendsOnly bool        `json:"pm_friends_only"`
	ProfileColour interface{} `json:"profile_colour"`
	Username      string      `json:"username"`
	CommentsCount int         `json:"comments_count"`
	CoverURL      string      `json:"cover_url"`
	Discord       interface{} `json:"discord"`
	HasSupported  bool        `json:"has_supported"`
	Interests     string      `json:"interests"`
	JoinDate      time.Time   `json:"join_date"`
	Kudosu        struct {
		Total     int `json:"total"`
		Available int `json:"available"`
	} `json:"kudosu"`
	Location     string      `json:"location"`
	MaxBlocks    int         `json:"max_blocks"`
	MaxFriends   int         `json:"max_friends"`
	Occupation   string      `json:"occupation"`
	Playmode     string      `json:"playmode"`
	Playstyle    []string    `json:"playstyle"`
	PostCount    int         `json:"post_count"`
	ProfileOrder []string    `json:"profile_order"`
	Title        interface{} `json:"title"`
	TitleURL     interface{} `json:"title_url"`
	Twitter      interface{} `json:"twitter"`
	Website      interface{} `json:"website"`
	Country      struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
	Cover struct {
		CustomURL string `json:"custom_url"`
		URL       string `json:"url"`
		ID        string `json:"id"`
	} `json:"cover"`
	AccountHistory         []interface{} `json:"account_history"`
	ActiveTournamentBanner interface{}   `json:"active_tournament_banner"`
	Badges                 []struct {
		AwardedAt   time.Time `json:"awarded_at"`
		Description string    `json:"description"`
		ImageURL    string    `json:"image_url"`
		URL         string    `json:"url"`
	} `json:"badges"`
	BeatmapPlaycountsCount   int           `json:"beatmap_playcounts_count"`
	FavouriteBeatmapsetCount int           `json:"favourite_beatmapset_count"`
	FollowerCount            int           `json:"follower_count"`
	GraveyardBeatmapsetCount int           `json:"graveyard_beatmapset_count"`
	Groups                   []interface{} `json:"groups"`
	LovedBeatmapsetCount     int           `json:"loved_beatmapset_count"`
	MappingFollowerCount     int           `json:"mapping_follower_count"`
	MonthlyPlaycounts        []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"monthly_playcounts"`
	Page struct {
		HTML string `json:"html"`
		Raw  string `json:"raw"`
	} `json:"page"`
	PreviousUsernames                []string `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int      `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCounts             []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"replays_watched_counts"`
	ScoresBestCount   int `json:"scores_best_count"`
	ScoresFirstCount  int `json:"scores_first_count"`
	ScoresRecentCount int `json:"scores_recent_count"`
	Statistics        struct {
		Level struct {
			Current  int `json:"current"`
			Progress int `json:"progress"`
		} `json:"level"`
		GlobalRank             int     `json:"global_rank"`
		Pp                     float64 `json:"pp"`
		RankedScore            int64   `json:"ranked_score"`
		HitAccuracy            float64 `json:"hit_accuracy"`
		PlayCount              int     `json:"play_count"`
		PlayTime               int     `json:"play_time"`
		TotalScore             int64   `json:"total_score"`
		TotalHits              int     `json:"total_hits"`
		MaximumCombo           int     `json:"maximum_combo"`
		ReplaysWatchedByOthers int     `json:"replays_watched_by_others"`
		IsRanked               bool    `json:"is_ranked"`
		GradeCounts            struct {
			Ss  int `json:"ss"`
			SSH int `json:"ssh"`
			S   int `json:"s"`
			Sh  int `json:"sh"`
			A   int `json:"a"`
		} `json:"grade_counts"`
		CountryRank int `json:"country_rank"`
		Rank        struct {
			Country int `json:"country"`
		} `json:"rank"`
	} `json:"statistics"`
	SupportLevel            int `json:"support_level"`
	UnrankedBeatmapsetCount int `json:"unranked_beatmapset_count"`
	UserAchievements        []struct {
		AchievedAt    time.Time `json:"achieved_at"`
		AchievementID int       `json:"achievement_id"`
	} `json:"user_achievements"`
	Rankhistory struct {
		Mode string `json:"mode"`
		Data []int  `json:"data"`
	} `json:"rankHistory"`
	RankHistory struct {
		Mode string `json:"mode"`
		Data []int  `json:"data"`
	} `json:"rank_history"`
}

type UserScores []struct {
	ID         int64    `json:"id"`
	UserID     int      `json:"user_id"`
	Accuracy   float64  `json:"accuracy"`
	Mods       []string `json:"mods"`
	Score      int      `json:"score"`
	MaxCombo   int      `json:"max_combo"`
	Perfect    bool     `json:"perfect"`
	Statistics struct {
		Count50   int `json:"count_50"`
		Count100  int `json:"count_100"`
		Count300  int `json:"count_300"`
		CountGeki int `json:"count_geki"`
		CountKatu int `json:"count_katu"`
		CountMiss int `json:"count_miss"`
	} `json:"statistics"`
	Rank      string    `json:"rank"`
	CreatedAt time.Time `json:"created_at"`
	BestID    int64     `json:"best_id"`
	Pp        float64   `json:"pp"`
	Mode      string    `json:"mode"`
	ModeInt   int       `json:"mode_int"`
	Replay    bool      `json:"replay"`
	Beatmap   struct {
		DifficultyRating float64     `json:"difficulty_rating"`
		ID               int         `json:"id"`
		Mode             string      `json:"mode"`
		Status           string      `json:"status"`
		TotalLength      int         `json:"total_length"`
		Version          string      `json:"version"`
		Accuracy         float64         `json:"accuracy"`
		Ar               float64        `json:"ar"`
		BeatmapsetID     int         `json:"beatmapset_id"`
		Bpm              int         `json:"bpm"`
		Convert          bool        `json:"convert"`
		CountCircles     int         `json:"count_circles"`
		CountSliders     int         `json:"count_sliders"`
		CountSpinners    int         `json:"count_spinners"`
		Cs               float64         `json:"cs"`
		DeletedAt        interface{} `json:"deleted_at"`
		Drain            float64        `json:"drain"`
		HitLength        int         `json:"hit_length"`
		IsScoreable      bool        `json:"is_scoreable"`
		LastUpdated      time.Time   `json:"last_updated"`
		ModeInt          int         `json:"mode_int"`
		Passcount        int         `json:"passcount"`
		Playcount        int         `json:"playcount"`
		Ranked           int         `json:"ranked"`
		URL              string      `json:"url"`
		Checksum         string      `json:"checksum"`
	} `json:"beatmap"`
	Beatmapset struct {
		Artist        string `json:"artist"`
		ArtistUnicode string `json:"artist_unicode"`
		Covers        struct {
			Cover       string `json:"cover"`
			Cover2X     string `json:"cover@2x"`
			Card        string `json:"card"`
			Card2X      string `json:"card@2x"`
			List        string `json:"list"`
			List2X      string `json:"list@2x"`
			Slimcover   string `json:"slimcover"`
			Slimcover2X string `json:"slimcover@2x"`
		} `json:"covers"`
		Creator        string      `json:"creator"`
		FavouriteCount int         `json:"favourite_count"`
		Hype           interface{} `json:"hype"`
		ID             int         `json:"id"`
		Nsfw           bool        `json:"nsfw"`
		PlayCount      int         `json:"play_count"`
		PreviewURL     string      `json:"preview_url"`
		Source         string      `json:"source"`
		Status         string      `json:"status"`
		Title          string      `json:"title"`
		TitleUnicode   string      `json:"title_unicode"`
		UserID         int         `json:"user_id"`
		Video          bool        `json:"video"`
	} `json:"beatmapset"`
	Weight struct {
		Percentage float64     `json:"percentage"`
		Pp         float64 `json:"pp"`
	} `json:"weight"`
	User struct {
		AvatarURL     string      `json:"avatar_url"`
		CountryCode   string      `json:"country_code"`
		DefaultGroup  string      `json:"default_group"`
		ID            int         `json:"id"`
		IsActive      bool        `json:"is_active"`
		IsBot         bool        `json:"is_bot"`
		IsDeleted     bool        `json:"is_deleted"`
		IsOnline      bool        `json:"is_online"`
		IsSupporter   bool        `json:"is_supporter"`
		LastVisit     interface{} `json:"last_visit"`
		PmFriendsOnly bool        `json:"pm_friends_only"`
		ProfileColour interface{} `json:"profile_colour"`
		Username      string      `json:"username"`
	} `json:"user"`
}
