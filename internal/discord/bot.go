package discord

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

var commandFuncs = []NewCommandFunc{
	NewMhwWeaknessCommand,
}

type Bot struct {
	session  *discordgo.Session
	commands []Command
}

func NewBot(c *config.Config) (*Bot, error) {
	token := fmt.Sprintf("Bot %s", c.DiscordBotAuthToken)
	s, err := discordgo.New(token)
	if err != nil {
		return nil, err
	}
	commands := make([]Command, 0)
	for _, newCommand := range commandFuncs {
		cmd := newCommand(c)
		commands = append(commands, cmd)
	}
	return &Bot{
		session:  s,
		commands: commands,
	}, nil
}

func (b *Bot) Setup(sh *search.SearchHandler) {
	b.session.Identify.Intents = discordgo.IntentsGuildMessages
	for _, cmd := range b.commands {
		b.session.AddHandler(cmd.Handle)
	}
}

func (b *Bot) withConnection(ctx context.Context, callback func(ctx context.Context) error) error {
	err := b.session.Open()
	if err != nil {
		return err
	}
	defer b.session.Close()
	if err = callback(ctx); err != nil {
		return err
	}
	return nil
}

func (b *Bot) Run() error {
	return b.withConnection(context.Background(), func(_ context.Context) error {
		fmt.Println("Bot is now running. Press CTRL-C to exit.")
		<-waitUntilCancelled()
		return nil
	})
}

func (b *Bot) CreateCommands(ctx context.Context, guildID string) error {
	return b.withConnection(ctx, func(ctx context.Context) error {
		for _, cmd := range b.commands {
			err := cmd.Create(ctx, b.session, guildID)
			if err != nil {
				return errors.Join(err, errors.New("unable to create command"))
			}
		}
		return nil
	})
}

func waitUntilCancelled() chan os.Signal {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	return sc
}
