package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, l := range links {

			fmt.Printf("%s\n", l)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %v\n", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Processed: %v\n", r.Request.URL)
	})

	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}