package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Gist json structure
type Gist struct {
	Rawurl string `json:"html_url"`
}

// Doer requests Gist API
type Doer interface {
	doGistsRequest(user string) (io.Reader, error)
}

type Client struct {
	Gister Doer
}

type Gister struct{}

func (g *Gister) doGistsRequest(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return &buf, nil
}

// ListGists return url list
func (c *Client) ListGists(user string) ([]string, error) {
	r, err := c.Gister.doGistsRequest(user)
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}

	return urls, nil
}

func main() {
	c := &Client{&Gister{}}
	urls, err := c.ListGists("DaisukeHirata")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(urls)
}
