package main

import (
	"github.com/go-bongo/bongo/makehttp"
)

func main()  {
	makehttp.Scraper("https://www.bjp-online.com/2019/10/portrait-of-humanity-the-undiscovered-work-of-michel-kameni/", "#primary article", "article")
}