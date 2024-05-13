package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jgabriel1/mh-weakness-bot/internal/config"
)

type Command interface {
	Create(s *discordgo.Session, guildID string) error
	Handle(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type NewCommandFunc func(c *config.Config) Command
