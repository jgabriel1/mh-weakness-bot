package main

import (
	"context"
	"flag"
	"log"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	errChan := make(chan error, 1)

	go func() {
		if err = bot.CreateCommands(ctx, guildID); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Fatalf("context cancelled: %v", ctx.Err())
	case err := <-errChan:
		log.Fatalf("error creating command: %v", err)
	}
}
