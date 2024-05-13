package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/sxmbaka/lacuna/mode"
)

type WordList []Word

type Word struct {
	Word  string   `json:"word"`
	Score int      `json:"score"`
	Tags  []string `json:"tags"`
}

const (
	baseURL = "https://api.datamuse.com/words"
)

func GetData(moreLikeActive *mode.MoreLike, soundsLikeActive *mode.SoundsLike, spelledLikeActive *mode.SpelledLike, max int, show bool) WordList {
	url := frameRequest(moreLikeActive, soundsLikeActive, spelledLikeActive, max, show)
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status: %s", response.Status)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response data: %v", err)
	}

	var wordlist WordList
	err = json.Unmarshal(responseData, &wordlist)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return wordlist
}
