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
	"https://www.washingtonpost.com/",
}

func main() {
	bench := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(urls))

	proccessed := new(uint8)
	for i := 0; i < len(urls); i++ {
		go func(i int) {
			parseSEO(proccessed, urls[i])
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}

func parseSEO(i *uint8, url string) {
	crawledPage := parseMeta(url)

	fmt.Println(crawledPage.format())
	*i += 1
	fmt.Printf("PROCCESSED %v/%v\n", *i, len(urls))
}
