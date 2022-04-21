package main

import (
	"encoding/json"
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

	seoData := []webPage{}

	proccessed := 0
	for _, url := range urls {
		go func(url string) {
			seoData = append(seoData, parseSEO(url))

			proccessed += 1
			fmt.Printf("PROCCESSED %v/%v\n", proccessed, len(urls))

			wg.Done()
		}(url)
	}

	wg.Wait()

	// fmt.Println(seoData)
	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}

func parseSEO(url string) webPage {
	crawledPage := parseMeta(url)

	data, err := json.MarshalIndent(crawledPage, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	return crawledPage
}
