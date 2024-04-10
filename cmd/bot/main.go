package main

import (
	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/discord"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

func main() {
	config, err := config.NewConfig("config")
	if err != nil {
		panic(err)
	}
	bot, err := discord.NewBot(config)
	if err != nil {
		panic(err)
	}
	searchHandler := search.NewSearchHandler(config)
	bot.Setup(searchHandler)
	if err = bot.Run(); err != nil {
		panic(err)
	}
}
