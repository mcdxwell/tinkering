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

	if *date == "" {
		showTitles(wotdURL)
	}

	if *date != "" {
		//date := *date
		fmt.Println("Date func")
	}
}

func HandleRnd(rndCmd *flag.FlagSet) {
	rndCmd.Parse(os.Args[2:])

	showTitles(getDateURL())
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

	// TODO: Make this a function that can be used to
	// append the wotd and random words to the json file.
	datey := word{
		Date: randomDate,
	}

	dateys := getWord()
	dateys = append(dateys, datey)
	saveWordo(dateys)
	return randomDate

}

// Concatenates the wotd URL with a random date
func getDateURL() string {

	var url strings.Builder
	url.WriteString(wotdURL)
	url.WriteString(getRandomDate())
	fmt.Println(url.String())

	return url.String()
}

func showTitles(url string) {

	c := getTitleTags(url)
	wordy := make([]string, 0)
	for msg := range c {
		wordy = append(wordy, msg)
	}

	fmt.Println(wordy[0])

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
