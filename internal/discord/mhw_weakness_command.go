package discord

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jgabriel1/mh-weakness-bot/internal/config"
	"github.com/jgabriel1/mh-weakness-bot/internal/scraping"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

const (
	mhwWeaknessCommandName        = "weaknessmhw"
	mhwWeaknessCommandDescription = "Get the elemental weakness for a monster in MH World, the best game in the franchise."
)

type MhwWeaknessCommand struct {
	searchHandler *search.SearchHandler
}

func NewMhwWeaknessCommand(c *config.Config) Command {
	sh := search.NewSearchHandler(c)
	return &MhwWeaknessCommand{searchHandler: sh}
}

func (cmd *MhwWeaknessCommand) Create(ctx context.Context, s *discordgo.Session, guildID string) error {
	_, err := s.ApplicationCommandCreate(s.State.Application.ID, guildID, &discordgo.ApplicationCommand{
		Name:        mhwWeaknessCommandName,
		Description: mhwWeaknessCommandDescription,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "name",
				Description: "Monster name to be searched.",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	})
	return err
}

// TODO add better error handling for command handlers
func (cmd *MhwWeaknessCommand) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}
	if i.ApplicationCommandData().Name != mhwWeaknessCommandName || len(i.ApplicationCommandData().Options) < 1 {
		return
	}
	nameOption := i.ApplicationCommandData().Options[0]
	name := strings.Trim(nameOption.StringValue(), " ")
	results, err := cmd.searchHandler.SearchMonsterName(name)
	if err != nil {
		respondInteraction("An unexpected error occured. Try again later.", s, i)
		return
	}
	if len(results) < 1 {
		respondInteraction(fmt.Sprintf("No monsters were found with name: \"%s\"", name), s, i)
		return
	}
	t, err := scraping.ScrapeMonsterHitzonesTable(results[0].URL)
	if err != nil {
		respondInteraction("An unexpected error occured. Try again later.", s, i)
		return
	}
	el, _ := t.GetWeaknessElement()
	respondInteraction(fmt.Sprintf("Weakness for %s is %s", results[0].Name, el), s, i)
}

func respondInteraction(message string, s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
}
