package seo

import "fmt"

type meta struct {
	Property string `json:"property"`
	Content  string `json:"content"`
}

type webPage struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	MetaTags []meta `json:"metaTags"`
}

func (w *webPage) format() string {
	str := fmt.Sprintf("")

	str += fmt.Sprintf("TITLE: %v\n", w.Title)
	str += fmt.Sprintln("---")
	str += fmt.Sprintf("URL: %v\n", w.Url)
	str += fmt.Sprintln("---")

	for _, v := range w.MetaTags {
		str += fmt.Sprintf("%v - %v\n", v.Property, v.Content)
		str += fmt.Sprintln("---")
	}

	return str
}
