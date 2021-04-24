package osu

import (
	"fmt"
	oppai "github.com/flesnuk/oppai5"
	"os"
)

type Performance struct {
	Pp          *oppai.PP
	BeatmapInfo *oppai.Map
	PpFc        *oppai.PP
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

		// calc pp if fc
		// turn misses into 300's
		parameters.N300 += parameters.Misses
		parameters.Misses = 0
		ppFC := pp
		ppFC.PP.PPv2WithMods(ppFC.Diff.Aim, ppFC.Diff.Speed, beatmap, int(parameters.Mods), int(parameters.N300), int(parameters.N100), int(parameters.N50), int(parameters.Misses), beatmap.MaxCombo)

		return &Performance{Pp: &pp, BeatmapInfo: beatmap, PpFc: &ppFC}, nil
	}

	return nil, fmt.Errorf("oppai returned nil map file")
}
