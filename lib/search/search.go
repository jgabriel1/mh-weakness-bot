package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jgabriel1/mh-weakness-bot/lib/config"
)

const searchUrl string = "https://search.kiranico.gg/indexes/mhworld_en_docsearch/search"

type SearchResponse struct {
	Hits []SearchResult `json:"hits"`
}

type SearchResult struct {
	Type string `json:"lvl0"`
	Name string `json:"lvl2"`
	URL  string `json:"url"`
}

func SearchMonsterName(q string) ([]SearchResult, error) {
	searchAPIAuthKey, err := config.GetSearchAPIAuthKey()
	if err != nil {
		return []SearchResult{}, err
	}

	reqBody := fmt.Sprintf(`{"q":"%s"}`, q)

	req, _ := http.NewRequest("POST", searchUrl, bytes.NewBufferString(reqBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", searchAPIAuthKey))

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
