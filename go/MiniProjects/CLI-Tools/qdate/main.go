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
	getMonths()
	getYears()
}

// return type: time.Time
func getCurrDate() time.Time {
	current_date := time.Now().Local()
	fmt.Println("The current date is: ", current_date.Format("01-02-2006"))
	return current_date
}

// Do the same with getMonth() and getYear()
func getDays() {
	days := flag.String("d", "d22", "Measures of time: day, month, year. Format: dX where X is the number of days")
	dd := getConv(*&days)
	desired_dd := getCurrDate().AddDate(0, 0, dd)
	fmt.Printf("The date in %d days: %v", dd, desired_dd.Format("01-02-2006"))
}

func getMonths() {
	months := flag.String("m", "m22", "Measures of time: day, month, year. Format: mX where X is the number of months")
	mm := getConv(*&months)
	desired_mm := getCurrDate().AddDate(0, mm, 0)
	fmt.Printf("The date in %d months: %v", mm, desired_mm.Format("01-02-2006"))

}

func getYears() {
	years := flag.String("y", "y22", "Measures of time: day, month, year. Format: yX where X is the number of years")
	yyyy := getConv(*&years)
	desired_yyyy := getCurrDate().AddDate(yyyy, 0, 0)
	fmt.Printf("The date in %d years: %v", yyyy, desired_yyyy.Format("01-02-2006"))
}

func getConv(toonies *string) int {

	tUnits := (*toonies)[1:]
	t, err := strconv.Atoi(tUnits)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
