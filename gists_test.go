package main

import (
	"io"
	"strings"
	"testing"
)

type dummyDoer struct{}

func (d *dummyDoer) doGistsRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
		[
			{"html_url": "https://gist.github.com/1a861e999f3c73f83146"},
			{"html_url": "https://gist.github.com/653d65c17d8fe1d372af"}				
		]
		`), nil
}

func TestGetGists(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatalf("list gets caused error: %s", err)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
