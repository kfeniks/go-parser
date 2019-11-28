package main

import (
	"github.com/go-bongo/bongo/makehttp"
)

func main()  {
	makehttp.Scraper("https://www.bjp-online.com", "aside article h3.entry-title")
}