package osu

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func (c *Client) DownloadBeatmapFile(beatmapID string) (*os.File, error) {

	// Send request to get beatmap file
	resp, err := http.Get(fmt.Sprintf("https://osu.ppy.sh/osu/%s", beatmapID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Generate a random file name
	rand.Seed(time.Now().UnixNano())
	suffix := fmt.Sprintf("%d", rand.Intn(10000))
	fileName := fmt.Sprintf("%s%s", suffix, ".osu")

	// Create the file
	out, err := os.Create(fmt.Sprintf("pkg/osu/maps/%s", fileName))
	if err != nil {
		return nil, err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}
	return out, err

}