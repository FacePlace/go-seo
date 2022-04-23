package seo

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func parseMeta(url string) webPage {
	c := colly.NewCollector()

	metaTags := []meta{}
	var title string

	c.OnHTML("head", func(e *colly.HTMLElement) {
		title = e.ChildText("title")
		e.ForEach("meta", func(_ int, e *colly.HTMLElement) {
			property := e.Attr("property")
			if len(property) == 0 {
				property = e.Attr("name")
			}
			content := e.Attr("content")

			if len(property) > 0 && len(content) > 0 {
				metaTags = append(metaTags, meta{
					property,
					content,
				})
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Failed requesting URL: %v with respone: %v\n", r.Request.URL, r)
		fmt.Println("Error: ", err)
	})

	c.Visit(url)

	crawledPage := webPage{
		title,
		url,
		metaTags,
	}

	return crawledPage
}
