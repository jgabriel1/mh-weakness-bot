package main

import (
	"fmt"

	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/scraping"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

func main() {
	config, err := config.NewConfig("config")
	if err != nil {
		panic(err)
	}

	searchHandler := search.NewSearchHandler(config)
	results, err := searchHandler.SearchMonsterName("rajang")
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
