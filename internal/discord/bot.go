package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

type Bot struct {
	session *discordgo.Session
}

func NewBot(c *config.Config) (*Bot, error) {
	token := fmt.Sprintf("Bot %s", c.DiscordBotAuthToken)
	s, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	return &Bot{
		session: s,
	}, nil
}

func (b *Bot) Setup(sh *search.SearchHandler) {
	b.session.Identify.Intents = discordgo.IntentsGuildMessages
	b.session.AddHandler(makeWeaknessElementFromQueryHandler(sh))
}

func (b *Bot) Run() error {
	err := b.session.Open()
	if err != nil {
		return err
	}
	defer b.session.Close()
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-waitUntilCancelled()
	return err
}

func waitUntilCancelled() chan os.Signal {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	return sc
}
