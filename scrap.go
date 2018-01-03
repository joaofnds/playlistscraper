package playlistscraper

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"
)

var (
	// ErrEmptyPlaylistID is the error returned when a empty playlist id is received
	ErrEmptyPlaylistID = fmt.Errorf("empty playlist id")

	// ErrCollyScrapeFail is the error returned when a colly Visit resulted in an error
	ErrCollyScrapeFail = fmt.Errorf("colly failed to scrape")
)

// ScrapeVideoLinks gets a youtube playlist ID,
// scrape links from it's videos(up to 100 links),
// and them return them
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
		err = ErrCollyScrapeFail
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return links, nil
}
