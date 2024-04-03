package main

import (
	"fmt"

	"github.com/jgabriel1/mh-weakness-bot/lib/scraping"
	"github.com/jgabriel1/mh-weakness-bot/lib/search"
)

func main() {
	results, err := search.SearchMonsterName("rajang")
	if err != nil {
		panic(err)
	}

	t, err := scraping.ScrapeMonsterHitzonesTable(results[0].URL)
	if err != nil {
		panic(err)
	}

	for _, col := range t.Columns {
		fmt.Println(col)
	}
}
