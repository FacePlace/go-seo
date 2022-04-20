package main

import (
	"fmt"
	"time"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	start := time.Now()

	// os.WriteFile("out.csv", []byte{}, 600)
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://finance.yahoo.com/world-indices"},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {

		},
		// Exporters: []export.Exporter{&export.CSV{}},
	}).Start()

	fmt.Printf("Finished proccessing: %v", time.Since(start))
}
