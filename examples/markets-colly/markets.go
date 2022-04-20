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
		e.ForEach("li", func(i int, e *colly.HTMLElement) {
			fmt.Println(e)
		})
	})

	c.Visit("https://finance.yahoo.com/world-indices")

	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}
