package playlistscraper

import (
	"net/url"
	"testing"
)

func TestScrapeVideoLinks(t *testing.T) {
	tt := []struct {
		name string
		pID  string
		vIDs []string
		err  error
	}{
		{"empty playlist id", "", nil, ErrEmptyPlaylistID},
		{"wrong playlist id", "asdf", []string(nil), nil},
		{
			"correct playlist id",
			"PLxqB4IDcfCjEtldWKdYI1KJoa9ilbtwp_",
			[]string{"KwufrpEDF9M", "Tss5ztW4E3M", "e5RuGj0g1tk", "ogrEX1oBav4", "BYWOV6jOapk"},
			nil},
	}

	// TODO: find a way to make colly fail

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			links, err := ScrapeVideoLinks(tc.pID)
			for i := range links {
				u, err := url.ParseRequestURI(links[i])
				if err != nil {
					t.Fatalf("failed to parse uri string")
				}

				if u.Query().Get("v") != tc.vIDs[i] {
					t.Errorf("[%s] expected video id to be %v, got: n%v\n", tc.name, tc.vIDs[i], links[i])
				}
			}

			if err != tc.err {
				t.Errorf("[%s] expected err to be %v, got: %v\n", tc.name, tc.err, err)
			}
		})
	}
}
