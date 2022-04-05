package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	b := Contact{}
	saveContact(b)
	
}

// Functional programming inspiration
// https://www.youtube.com/watch?v=Up7LcbGZFuo
// "Persistence ignorance"

// Fun with composition
type FullName struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

type Addresses struct {
	Address         string `json:"address,omitempty"`
	Email           string `json:"email,omitempty"`
	VerifiedEmail   string `json:"verifiedemail,omitempty"`
	UnverifiedEmail string `json:"unverifiedemail,omitempty"`
}

type Contact struct {
	ID		  string     `json:"id"`
	FullName  *FullName  `json:"fullname,omitempty"`
	Addresses *Addresses `json:"addresses,omitempty"`
}

// Which values are optional?
// What is the domain logic?

type fn func(int)

func test(f fn, val int) {
	f(val)
}

func hashInfo(info string) string {
	hashedInfo := sha256.Sum256([]byte(info))
	var shash string = fmt.Sprintf("%x", hashedInfo)
	return shash
}

func resetPassword(a Addresses) {
	if hashInfo(a.Email) != a.VerifiedEmail {
		fmt.Println("Unavailable")
		return
	}
	fmt.Println("Available")
	return
}

func saveContact(c Contact) {

	// read and write word information to existing information in wotd.json
	info := getContacts()
	fmt.Println(info)
	info = append(info,
		Contact{
			ID: "321dfjds8j",
			FullName: &FullName{
				First: "Peggy", Middle: "Bigfoot", Last: "Hill"},
			Addresses: &Addresses{
				Email:         "peggy@peggy.com",
				VerifiedEmail: hashInfo("peggy@peggy.com"),
				Address:       "84 Rainey Street, Arlen, Texas"}})

	saveContactInfo(info)
}

func saveContactInfo(c []Contact) {
	wordData, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./stuff.json", wordData, 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func getContacts() (contact []Contact) {

	wordData, err := os.ReadFile("./stuff.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(wordData, &contact)

	if err != nil {
		log.Fatal(err)
	}

	return contact
}
