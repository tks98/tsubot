package osu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *Client) GetUser(id string) (*User, error) {

	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/users/%s/osu", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
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

	return &user, nil
}

func (c *Client) GetUserScores(id string, kind string, offset string) (*User, error) {
	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/users/%s/scores/kind?include_fails=1", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(bodyBytes))

	return nil, err

}


