package discord

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/jgabriel1/mh-weakness-bot/internal/scraping"
	"github.com/jgabriel1/mh-weakness-bot/internal/search"
)

func makeWeaknessElementFromQueryHandler(searchHandler *search.SearchHandler) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		query, ok := matchesWeaknessCommand(m.Content)
		if !ok {
			return
		}
		results, err := searchHandler.SearchMonsterName(query)
		if err != nil {
			replyMessage("An unexpected error occured. Try again later.", s, m)
			return
		}
		if len(results) < 1 {
			replyMessage(fmt.Sprintf("No monsters were found with name: \"%s\"", query), s, m)
			return
		}
		t, err := scraping.ScrapeMonsterHitzonesTable(results[0].URL)
		if err != nil {
			replyMessage("An unexpected error occured. Try again later.", s, m)
			return
		}
		el, _ := t.GetWeaknessElement()
		replyMessage(fmt.Sprintf("Weakness for %s is %s", results[0].Name, el), s, m)
	}
}

func replyMessage(message string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, message, m.Reference())
}

func matchesWeaknessCommand(q string) (string, bool) {
	re := regexp.MustCompile(`(?i)\.weakness\s+(.+)`)
	matches := re.FindStringSubmatch(q)
	if len(matches) < 2 {
		return "", false
	}
	return matches[1], true
}
