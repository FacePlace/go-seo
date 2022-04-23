package seo_test

import (
	"testing"

	"github.com/FacePlace/go-seo/seo"
)

func TestParseSEO(t *testing.T) {
	urls := []string{
		"https://finance.yahoo.com/world-indices",
		"https://www.cnn.com/",
		"https://www.nytimes.com/",
		"https://nypost.com/",
		"https://www.wsj.com/",
		"https://www.reuters.com/",
		"https://www.cbsnews.com/",
		"https://www.cnbc.com/world/",
		"https://www.npr.org/",
		// "https://www.washingtonpost.com/",
	}

	err := seo.GetSEO(urls)

	if err != nil {
		t.Fatalf("Failed test with error: %v", err)
	}
}
