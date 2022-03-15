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

// "Persistence ignorance"

type FullName struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

type Addresses struct {
	Address         string `json:"address,omitempty"`
	Email           string `json:"email,omitempty"`
	VerifiedEmail   string `json:"verifiedEmail,omitempty"`
	UnverifiedEmail string `json:"unverifiedEmail,omitempty"`
}

type ContactInfo struct {
	Addresses
}

type Contact struct {
	FullName
	ContactInfo
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

	c.First = "Did"
	c.Middle = ""
	c.Last = "McDowe"
	c.Email = "davd@gmail"
	c.VerifiedEmail = "fdssfd"
	//c.VerifiedEmail = hashInfo(c.Email)
	c.Address = "123 Fake St."
	//fmt.Println(hashInfo(c.Email))
	//fmt.Println(c)
	//resetPassword(c.Addresses)

	// read and write word information to existing information in wotd.json
	info := getContacts()
	info = append(info, c)
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
