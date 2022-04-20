package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

type meta struct {
	property string
	content  string
}

type webPage struct {
	title    string
	url      string
	metaTags []meta
}

func (w *webPage) format() string {
	str := fmt.Sprintf("")

	str += fmt.Sprintf("TITLE: %v\n", w.title)
	str += fmt.Sprintln("---")
	str += fmt.Sprintf("URL: %v\n", w.url)
	str += fmt.Sprintln("---")

	for _, v := range w.metaTags {
		str += fmt.Sprintf("%v - %v\n", v.property, v.content)
		str += fmt.Sprintln("---")
	}

	return str
}

func main() {
	bench := time.Now()

	c := colly.NewCollector()

	metaTags := []meta{}
	var title string
	url := "https://finance.yahoo.com/world-indices"

	c.OnHTML("head", func(e *colly.HTMLElement) {
		title = e.ChildText("title")
		e.ForEach("meta", func(i int, e *colly.HTMLElement) {
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

	c.Visit(url)

	crawledPage := webPage{
		title,
		url,
		metaTags,
	}

	fmt.Println(crawledPage.format())

	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}
