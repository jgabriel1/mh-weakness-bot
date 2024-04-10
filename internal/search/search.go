package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jgabriel1/mh-weakness-bot/internal/config"
)

const searchUrl string = "https://search.kiranico.gg/indexes/mhworld_en_docsearch/search"

type SearchHandler struct {
	config *config.Config
}

type SearchResponse struct {
	Hits []SearchResult `json:"hits"`
}

type SearchResult struct {
	Type string `json:"lvl0"`
	Name string `json:"lvl2"`
	URL  string `json:"url"`
}

func NewSearchHandler(config *config.Config) *SearchHandler {
	return &SearchHandler{
		config: config,
	}
}

func (s *SearchHandler) SearchMonsterName(q string) ([]SearchResult, error) {
	reqBody := fmt.Sprintf(`{"q":"%s"}`, q)

	req, err := http.NewRequest("POST", searchUrl, bytes.NewBufferString(reqBody))
	if err != nil {
		return []SearchResult{}, fmt.Errorf("failed to create http req = %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.config.SearchAPIKey))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return []SearchResult{}, errors.Join(
			errors.New("there was an error during the search request"), err)
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	parsedRes := &SearchResponse{}
	err = json.Unmarshal(resBody, parsedRes)
	if err != nil {
		return []SearchResult{}, errors.Join(
			errors.New("the search response format is not parseable"), err)
	}

	monsters := []SearchResult{}
	for _, hit := range parsedRes.Hits {
		if hit.Type == "Monsters" {
			monsters = append(monsters, hit)
		}
	}
	return monsters, nil
}
