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
	//getConv()
}

// return type: time.Time
func getCurrDate() time.Time {
	current_date := time.Now().Local()
	fmt.Println("The current date is: ", current_date.Format("01-02-2006"))
	return current_date
}

func getDays() {
	// TO-DO:
	// make lines 25 to 31 one function
	// ----
	// completed
	days := flag.String("d", "d22", "Measures of time: day, month, year. Format: dX where X is the number of days")

	dd := getConv(*&days)

	// TO-DO:
	// use AddDate() instead of Add(), cleaner and more consistent
	// https://pkg.go.dev/time#example-Time.Add
	// ----
	// completed
	desired_dd := getCurrDate().AddDate(0, 0, dd)
	// %d
	fmt.Printf("The date in %d days: %v", dd, desired_dd.Format("01-02-2006"))

}

func getMonths() {

}

func getYears() {

}

func getConv(toonies *string) int {

	tUnits := (*toonies)[1:]
	t, err := strconv.Atoi(tUnits)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
