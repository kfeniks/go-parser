package main

import (
	"github.com/go-bongo/bongo/makehttp"
	"strconv"
)

func main()  {
	for i := 1; i < 40; i++ {
		makehttp.Scraper("https://www.bjp-online.com/archive/page/" + strconv.Itoa(i), "aside article h3.entry-title")
	}
	//makehttp.Scraper("https://www.bjp-online.com", "aside article h3.entry-title")
}