package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://coinmarketcap.com/all/views/all/"},
		ParseFunc: coinsParse,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func coinsParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("tr.cmc-table-row").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"name": s.Find("a.cmc-table__column-name--name").Text(),
		}
	})
}
