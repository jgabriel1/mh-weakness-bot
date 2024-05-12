package main

import (
	"flag"

	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/discord"
)

// TODO: maybe make this a handler for when the bot gets added to a guild or figure out a better way to create commands
func main() {
	guildID := *flag.String("guildID", "", "Guild ID to create commands.")
	flag.Parse()
	c, err := config.NewConfig("config")
	if err != nil {
		panic("error loading config")
	}
	bot, err := discord.NewBot(c)
	if err != nil {
		panic(err)
	}
	if err = bot.CreateCommands(guildID); err != nil {
		panic(err)
	}
}
