package makehttp

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func SendRequest(client *http.Client, url string, selector string, typePage string) {
	doc, _ := getResponse(client, url)
	findSelector(doc, selector, typePage)
}

func getResponse(client *http.Client, url string) (*goquery.Document, error)  {
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

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc, err
}

func findSelector(doc *goquery.Document, selector string, typePage string) {
	// Find the article items
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {

		if typePage == "links" {
			// For each item found, get the link and title
			linkData := s.Find("a")
			link, _ := linkData.Attr("href")
			title := linkData.Text()
			fmt.Printf("Review %d: %s - %s\n", i, link, title)
		}

		if typePage == "article" {
			// For each item found, get the link and title
			articleData := s.Find(".entry-content")
			text := articleData.Text()
			fmt.Printf("%s\n", text)
		}
	})
}

func Scraper(url string, selector string, typePage string) {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	SendRequest(client, url, selector, typePage)
}