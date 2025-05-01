package jsonDecoder

import (
	"encoding/json"
	"os"
)

type StoryChoices struct {
	Intro struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"intro"`
	NewYork struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"new-york"`
	Debate struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"debate"`
	SeanKelly struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"sean-kelly"`
	MarkBates struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"mark-bates"`
	Denver struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	} `json:"denver"`
	Home struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"home"`
	}
}
type StoryChoice struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}

type storyMap map[string]StoryChoice

func GetOptionsFromJSON() (mappedOptions storyMap, err error) {

	var storyChoices StoryChoices
	file, err := os.ReadFile("../gopher.json")
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(file, &storyChoices)
	if jsonErr != nil {
		return nil, err
	}

	mappedOptions = map[string]StoryChoice{
		"intro":      StoryChoice(storyChoices.Intro),
		"new-york":   StoryChoice(storyChoices.NewYork),
		"debate":     StoryChoice(storyChoices.Debate),
		"sean-kelly": StoryChoice(storyChoices.SeanKelly),
		"mark-bates": StoryChoice(storyChoices.MarkBates),
		"denver":     StoryChoice(storyChoices.Denver),
		"home":       StoryChoice(storyChoices.Home),
	}
	return mappedOptions, nil
}
