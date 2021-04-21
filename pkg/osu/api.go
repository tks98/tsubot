package osu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) GetUserGlobalRank(id string) (int, error) {

	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/users/%s/osu", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user.Statistics.GlobalRank, nil
}
