package playlistscraper

import (
	"net/url"
	"reflect"
	"testing"
)

func TestScrapeVideoLinks(t *testing.T) {
	_, err := ScrapeVideoLinks("")
	if err != ErrEmptyPlaylistID {
		t.Errorf("Expect 'Empty Playlist ID', but got: %v\n", err)
	}

	links, err := ScrapeVideoLinks("PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ")
	if err != nil {
		t.Errorf("Expect correct PlaylistID scrape to not return an error, but got: %v\n", err)
	}

	if reflect.TypeOf(links) != reflect.SliceOf(reflect.TypeOf("")) {
		t.Errorf("Expect links to be of type []string, but got: %T\n", links)
	}

	for _, l := range links {
		if _, err = url.ParseRequestURI(l); err != nil {
			t.Errorf("Expect link to be a valid URI, got: %v\n", l)
		}
	}
}
