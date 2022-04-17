package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	os.WriteFile("out.csv", []byte{}, 600)
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://coinmarketcap.com/all/views/all/"},
		ParseFunc: coinsParse,
		Exporters: []export.Exporter{&export.CSV{}},
	}).Start()
}

func coinsParse(g *geziyor.Geziyor, r *client.Response) {
	coins := r.HTMLDoc.Find("tr.cmc-table-row")
	coins.Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"name":  fmt.Sprintf(s.Find("a.cmc-table__column-name--name").Text()),
			"price": fmt.Sprintf("$%v", s.Find("a.cmc-link").Text()),
		}
	})
}
