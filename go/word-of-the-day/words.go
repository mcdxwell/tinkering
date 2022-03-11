package main

import (
	"encoding/json"
	"os"
)

type word struct {
	Date       string `json:"date"`
	Word       string `json:"word"`
	Class      string `json:"class,omitempty"` // major form classes (noun, verb, adjective, and adverb)
	Meaning    string `json:"meaning,omitempty"`
	Definition string `json:"definition,omitempty"`
	Example    string `json:"example,omitempty"`
}

func getWord() (wotd []word) {
	wordData, err := os.ReadFile("./wotd.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(wordData, &wotd)

	if err != nil {
		panic(err)
	}

	return wotd
}

func saveWord(w []word) {
	wordData, err := json.Marshal(w)

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./wotd.json", wordData, 0644)

	if err != nil {
		panic(err)
	}
}
