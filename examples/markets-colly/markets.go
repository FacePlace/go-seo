package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	bench := time.Now()

	c := colly.NewCollector()

	c.OnHTML("ul.Carousel-Slider", func(e *colly.HTMLElement) {
		for _, v := range e.ChildTexts("li") {
			fmt.Println(v)
		}
	})

	c.Visit("https://finance.yahoo.com/world-indices")

	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}
