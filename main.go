package main

import (
	"fmt"
	"time"
)

func main() {
	bench := time.Now()

	urls := []string{
		"https://finance.yahoo.com/world-indices",
		"https://www.cnn.com/",
		"https://www.nytimes.com/",
		"https://nypost.com/",
		"https://www.washingtonpost.com/",
	}

	for i, v := range urls {
		crawledPage := parseMeta(v)

		fmt.Println(crawledPage.format())
		fmt.Printf("PROCCESSED %v/%v\n", i+1, len(urls))
	}

	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}
