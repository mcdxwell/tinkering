package main

import (
	"encoding/json"
	"io/ioutil"
)

type word struct {
	Word       string
	Class      string // major form classes (noun, verb, adjective, and adverb)
	Meaning    string
	Definition string
	Example    string
	Date       string
}

func getWord() {

}

func saveWord(w []word) {
	wordData, err := json.Marshal(w)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./words.json", wordData, 0644)

	if err != nil {
		panic(err)
	}
}
