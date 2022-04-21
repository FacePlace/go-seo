package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	f, err := os.OpenFile("seo.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	err = f.Truncate(0)
	_, err = f.Seek(0, 0)

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

	data, err := json.MarshalIndent(seoData, "", "  ")

	if err != nil {
		panic(err)
	}

	f.WriteString(string(data))

	// fmt.Println(seoData)
	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}

func parseSEO(url string) webPage {
	crawledPage := parseMeta(url)

	return crawledPage
}
