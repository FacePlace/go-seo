package main

import "fmt"

type meta struct {
	property string
	content  string
}

type webPage struct {
	title    string
	url      string
	metaTags []meta
}

func (w *webPage) format() string {
	str := fmt.Sprintf("")

	str += fmt.Sprintf("TITLE: %v\n", w.title)
	str += fmt.Sprintln("---")
	str += fmt.Sprintf("URL: %v\n", w.url)
	str += fmt.Sprintln("---")

	for _, v := range w.metaTags {
		str += fmt.Sprintf("%v - %v\n", v.property, v.content)
		str += fmt.Sprintln("---")
	}

	return str
}
