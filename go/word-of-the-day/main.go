package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var wg sync.WaitGroup

const wotdURL = "https://www.merriam-webster.com/word-of-the-day/"

// Everything needs to refactored aside from a couple of functions
// Use cases for wotd
// 1. Getting the word of the day
// 2. Generating a random word of the day
// 3. Generating a random word of the day by providing a date (optional - not implemented)

// Events that must take place to present a word:
// 1. The user picks between 3 commands - `get` the wotd, get `random` word, random word given a `date` (3rd is optional).
// 2. The program then generates the current date.
// 3. If the current date or the random date are present in the json file => return word info.
// 4. Else, fetch the word given the date and its information.
// 5. Save the word and its information to wotd.json

// TODO: Generate the current date in the wotd function
// If the date already exists in the json file, fetch word info from the json file. Else, fetch from the webpage.
func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError) //today
	//getToday := getCmd.Bool("today", false, "Get today's WOTD")
	//getRandom := getCmd.Bool("random", false, "Get random word")
	rndCmd := flag.NewFlagSet("random", flag.ExitOnError)
	getDate := getCmd.String("date", "", "Get word on this date")

	if len(os.Args) < 2 {
		fmt.Println("expected `get` subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getDate)
	case "random":
		HandleRnd(rndCmd)
	default:
	}

	//getDateURL()
	//showTitles(wotdURL)
	//fmt.Println(getDateURL())
}

func HandleGet(getCmd *flag.FlagSet, date *string) {

	getCmd.Parse(os.Args[2:])
	today := time.Now().UTC()
	if *date == "" {
		fmt.Println(today.Format("2006-01-02"))
		checkDate(today.Format("2006-01-02"))
	}

	if *date != "" {
		//date := *date
		fmt.Println("Date func")
	}
}

func HandleRnd(rndCmd *flag.FlagSet) {
	rndCmd.Parse(os.Args[2:])
	date := getRandomDate()
	checkDate(date)
}

// Creates a random date between 2006-10-10 to the current date
// When you pass a negative integer to AddDate(), you get the date for (x) days ago.
// FYI: This technique could be used to fetch a date for (y) months or years ago.
func getRandomDate() string {

	start := time.Date(2006, 10, 10, 0, 0, 0, 0, time.UTC)
	today := time.Now().UTC()
	diff := today.Sub(start)
	d := int(diff.Hours() / 24)
	rand.Seed(time.Now().UnixNano())
	randomDay := rand.Intn(d - 0)
	randomDate := today.AddDate(0, 0, -randomDay).Format("2006-01-02")

	return randomDate

}

// Use the date of a word to reference the json file.
// This eliminates the possibility of having more than 1 word
// per day, but it does not eliminate reused words.
// FYI: Merriam-Webster tends to reuse words.
func checkDate(d string) string {

	words := getWord()

	for _, word := range words {
		fmt.Println(word.Date)
		if word.Date == d {
			fmt.Println("From json:", word.Date, word.Word)
			return word.Word // return the word - or the word struct with all the word information
			// I have yet to decide whether or not I want to return just the word or word w/ information
		}

	}
	w := showTitles(getDateURL(d))

	fmt.Println("Not stored in json: ", w)
	saveInfo(d, w)
	return w
}

func formatter(title string) string {
	res := strings.ReplaceAll(title, "Word of the Day: ", "")
	res = strings.ReplaceAll(res, " | Merriam-Webster", "")
	fmt.Println(res)
	return res
}

// params:
// date, word, class, meaning, defintion, and example
func saveInfo(d, w string) {

	// TODO: Make this a function that can be used to
	// append the wotd and random words to the json file.
	wInfo := word{
		Date: d,
		Word: w,
	}

	// read and write word information to existing information in wotd.json
	info := getWord()
	info = append(info, wInfo)
	saveWord(info)
}

// Concatenates the wotd URL with a random date
func getDateURL(date string) string {

	var url strings.Builder
	url.WriteString(wotdURL)
	url.WriteString(date)
	//fmt.Println(url.String())

	return url.String()
}

func showTitles(url string) string {

	c := getTitleTags(url)
	wordy := make([]string, 0)
	for msg := range c {
		wordy = append(wordy, msg)
	}

	return formatter(wordy[0])

}

func getTitleTags(url string) chan string {

	c := make(chan string)

	wg.Add(1)
	go getTitle(url, c)

	go func() {
		wg.Wait()

		close(c)
	}()

	return c
}

// Helpful reference
// https://zetcode.com/golang/net-html/
func getTitle(url string, c chan string) {

	fmt.Println("Url", url)
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		c <- "failed to fetch data"
		return
	}

	defer res.Body.Close()

	tkn := html.NewTokenizer(res.Body)

	var isTitle bool

	for {

		tt := tkn.Next()

		switch {
		case tt == html.ErrorToken:
			return

		case tt == html.StartTagToken:

			t := tkn.Token()

			isTitle = t.Data == "title"

		case tt == html.TextToken:

			t := tkn.Token()

			if isTitle {

				c <- t.Data
				isTitle = false
			}
		}
	}
}
