package makehttp

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func SendRequest(client *http.Client, url string, selector string) {
	// Create and modify HTTP request before sending
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")

	// Make request
	response, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()



	if response.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	findSelector(response, selector);
}

func findSelector(response *http.Response, selector string) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the article items
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the link and title
		linkData := s.Find("a")
		link, _ := linkData.Attr("href")
		title := linkData.Text()
		fmt.Printf("Review %d: %s - %s\n", i, link, title)
	})
}

func Scraper(url string, selector string) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	SendRequest(client, url, selector)
}