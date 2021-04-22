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

func (c *Client) GetUserScores(id string, kind string, offset string) (*UserScores, error) {
	id, err := c.GetUserID(id)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://osu.ppy.sh/api/v2/users/%s/scores/%s/", id, kind)
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

	var scores UserScores
	err = json.Unmarshal(bodyBytes, &scores)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(bodyBytes))

	return &scores, nil

}

func (c *Client) GetUserID(name string) (string, error) {
	user, err := c.GetUser(name)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", user.ID), nil
}
