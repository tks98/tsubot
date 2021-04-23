package osu

import (
	"fmt"
	oppai "github.com/flesnuk/oppai5"
	"os"
)

func (c *Client) PerformanceCalc(beatmapFile *os.File, parameters *oppai.Parameters) (*oppai.PP, error) {

	defer os.Remove(beatmapFile.Name())

	f, err := os.Open(beatmapFile.Name())
	if err != nil {
		return nil, err
	}

	beatmap := oppai.Parse(f)

	if beatmap != nil {
		result := oppai.PPInfo(beatmap, parameters)
		return &result, nil
	}





	return nil, fmt.Errorf("oppai returned nil map file")
}
