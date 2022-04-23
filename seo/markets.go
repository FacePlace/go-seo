package seo

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func GetSEO(urls []string) (string, error) {
	bench := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	f, err := os.OpenFile("seo.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}

	err = f.Truncate(0)
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}

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

	return string(data), err
}

func parseSEO(url string) webPage {
	crawledPage := parseMeta(url)

	return crawledPage
}
