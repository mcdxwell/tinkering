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

func main() {

	//getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	//getToday := getCmd.Bool("today", false, "Get today's WOTD")
	//getRandom := getCmd.Bool("random", false, "Get random word")
	//getDate := getCmd.String("date", "", "Get word on this date")

	/* if len(os.Args) < 2 {
		fmt.Println("expected `get` subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet()
	} */
	/* urls := []string{
		"https://www.merriam-webster.com/word-of-the-day",
		"https://www.merriam-webster.com/word-of-the-day/2021-10-10",
	} */
	getDateURL()
	//showTitles(urls)
	//fmt.Println(getDateURL())
}

func HandleGet(getCmd *flag.FlagSet, today *bool, random *bool, date *string) {

	getCmd.Parse(os.Args[2:])

	if true {

	}
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

// Concatenates the wotd URL with a random date
func getDateURL() string {

	var url strings.Builder
	url.WriteString("https://www.merriam-webster.com/word-of-the-day/")
	url.WriteString(getRandomDate())
	fmt.Println(url.String())

	return url.String()
}

func showTitles(urls []string) {

	c := getTitleTags(urls)

	for msg := range c {

		fmt.Println(msg)
	}
}

func getTitleTags(urls []string) chan string {

	c := make(chan string)

	for _, url := range urls {
		wg.Add(1)
		go getTitle(url, c)
	}

	go func() {
		wg.Wait()

		close(c)
	}()

	return c
}

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
