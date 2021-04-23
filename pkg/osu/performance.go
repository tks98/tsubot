package osu

import (
	"fmt"
	oppai "github.com/flesnuk/oppai5"
	"os"
)

type Performance struct {
	Pp          *oppai.PP
	BeatmapInfo *oppai.Map
}

func (c *Client) PerformanceCalc(beatmapFile *os.File, parameters *oppai.Parameters) (*Performance, error) {

	defer os.Remove(beatmapFile.Name())

	f, err := os.Open(beatmapFile.Name())
	if err != nil {
		return nil, err
	}

	beatmap := oppai.Parse(f)
	if beatmap != nil {
		pp := oppai.PPInfo(beatmap, parameters)
		return &Performance{Pp: &pp, BeatmapInfo: beatmap}, nil
	}





	return nil, fmt.Errorf("oppai returned nil map file")
}
