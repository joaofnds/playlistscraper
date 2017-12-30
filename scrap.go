package playlistscraper

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
)

var (
	// ErrEmptyPlaylistID ...
	ErrEmptyPlaylistID = fmt.Errorf("Empty Playlist ID")
	// ErrCollyScrapeFail ...
	ErrCollyScrapeFail = fmt.Errorf("Colly failed to scrape")
)

// ScrapeVideoLinks gets a youtube playlist ID, scrape links from it's videos(up to 100 links), and them return them
func ScrapeVideoLinks(pID string) ([]string, error) {
	if pID == "" {
		return nil, ErrEmptyPlaylistID
	}
	c := colly.NewCollector()

	var links []string
	c.OnHTML("a[href].pl-video-title-link", func(e *colly.HTMLElement) {
		l := fmt.Sprintf("https://youtube.com%s", e.Attr("href"))
		links = append(links, l)
	})

	var wg sync.WaitGroup
	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	var err error
	c.OnError(func(c *colly.Response, e error) {
		err = ErrCollyScrapeFail
		wg.Done()
	})

	wg.Add(1)
	if c.Visit("https://www.youtube.com/playlist?list="+pID) != nil {
		log.Fatal(ErrCollyScrapeFail)
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return links, nil
}
