package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	getDays()
}

// return type: time.Time
func getCurrDate() time.Time {
	current_date := time.Now().Local()
	fmt.Println("The current date is: ", current_date.Format("01-02-2006"))
	return current_date
}

func getDays() {
	// work in progress
	days := flag.String("d", "d20", "Measures of time: day, month, year. Format: dX where X is the number of days")

	numDays := (*days)[1:]
	d, err := strconv.Atoi(numDays)
	if err != nil {
		log.Fatal(err)
	}

	//current_date.Add(time.Hour * 24 * time.Duration(d))
	desired_date := getCurrDate().Add(time.Hour * 24 * time.Duration(d))
	// %d
	fmt.Println("The date in days: ", desired_date.Format("01-02-2006"))

}

func getMonths() {

}
