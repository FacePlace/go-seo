package main

import (
	"fmt"
	"time"
)

func main() {
	bench := time.Now()

	url := "https://finance.yahoo.com/world-indices"

	crawledPage := parseMeta(url)

	fmt.Println(crawledPage.format())

	fmt.Printf("Finished proccessing: %v", time.Since(bench))
}
