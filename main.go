package main

import (
	"fmt"
	"sync"
	"time"
)

var urls = []string{
	"https://finance.yahoo.com/world-indices",
	"https://www.cnn.com/",
	"https://www.nytimes.com/",
	"https://nypost.com/",
	"https://www.wsj.com/",
	"https://www.reuters.com/",
	"https://www.cbsnews.com/",
	"https://www.cnbc.com/world/",
	"https://www.npr.org/",
	// "https://www.washingtonpost.com/",
}

func main() {
	bench := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(urls))

	proccessed := 0
	for i, url := range urls {
		go func(i int, url string) {
			parseSEO(url)
			proccessed += 1
			fmt.Printf("PROCCESSED %v/%v\n", proccessed, len(urls))
			wg.Done()
		}(i, url)
	}

	wg.Wait()
	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}

func parseSEO(url string) {
	crawledPage := parseMeta(url)

	fmt.Println(crawledPage.format())
}
