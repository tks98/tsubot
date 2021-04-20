package osu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	oauth  OauthResp
}

type OauthResp struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

// NewClient performs the oauth via client creds grant
func NewClient(clientID, clientSecret string) (*Client, error) {

	var api Client
	requestURL := fmt.Sprintf("https://osu.ppy.sh/oauth/token")
	v := url.Values{}
	v.Set("client_id", clientID)
	v.Set("client_secret", clientSecret)
	v.Set("grant_type", "client_credentials")
	v.Set("scope", "public")

	resp, err := http.PostForm(requestURL, v)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body OauthResp
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, err
	}

	// test api usage
	api.client = &http.Client{}
	api.oauth = body
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/v2/users/124493/osu", nil)
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + api.oauth.AccessToken
	req.Header.Add("Authorization", bearer)

	resp, err = api.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("osu api request test failed with response %s", resp.Status)
	}

	return &api, nil

}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	bearer := "Bearer " + c.oauth.AccessToken
	req.Header.Add("Authorization", bearer)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
