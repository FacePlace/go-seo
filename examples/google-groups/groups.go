package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Mail struct {
	Title   string
	Link    string
	Author  string
	Date    string
	Message string
}

func main() {
	var groupName string
	flag.StringVar(&groupName, "group", "hspbp", "Google Groups group name")
	flag.Parse()

	threads := make(map[string][]Mail)

	threadCollector := colly.NewCollector()
	mailCollector := colly.NewCollector()

	threadCollector.OnHTML("tr", func(e *colly.HTMLElement) {
		ch := e.DOM.Children()
		author := ch.Eq(1).Text()

		if author == "" {
			return
		}

		title := ch.Eq(0).Text()
		link, _ := ch.Eq(0).Children().Eq(0).Attr("href")

		link = strings.Replace(link, ".com/d/topic", ".com/forum/?_escaped_fragment_=topic", 1)
		date := ch.Eq(2).Text()

		log.Printf("Thread found: %s %q %s %s\n", link, title, author, date)
		mailCollector.Visit(link)
	})

	threadCollector.OnHTML("body > a[href]", func(e *colly.HTMLElement) {
		log.Println("Next page link found: ", e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	mailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		threadSubject := e.ChildText("h2")
		if _, ok := threads[threadSubject]; !ok {
			threads[threadSubject] = make([]Mail, 0, 8)
		}

		e.ForEach("table tr", func(_ int, el *colly.HTMLElement) {
			mail := Mail{
				Title:   el.ChildText("td:nth-of-type(1)"),
				Link:    el.ChildAttr("td:nth-of-type(1)", "href"),
				Author:  el.ChildText("td:nth-of-type(1)"),
				Date:    el.ChildText("td:nth-of-type(1)"),
				Message: el.ChildText("td:nth-of-type(1)"),
			}

			threads[threadSubject] = append(threads[threadSubject], mail)
		})

		if link, found := e.DOM.Find("> a[href]").Attr("href"); found {
			e.Request.Visit(link)
		} else {
			log.Printf("Thread %q done\n", threadSubject)
		}
	})

	threadCollector.Visit("https://groups.google.com/forum/?_escaped_fragment_=forum/" + groupName)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	enc.Encode(threads)
}
