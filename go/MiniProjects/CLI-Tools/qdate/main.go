package main

import (
	"fmt"
	"time"
)

func main() {
	current_date := time.Now().Local()
	fmt.Println("The current date is: ", current_date.Format("01-02-2006"))

	desired_date := current_date.Add(time.Hour * 24 * 10)
	// %d
	fmt.Println("The date in days: ", desired_date.Format("01-02-2006"))
}

func getDays() {
	// work in progress
	/* days := flag.String("d", "d6", "Measures of time: day, month, year. Format: dX where X is the number of days")

	numDays := (*days)[1:]
	d, err := strconv.Atoi(numDays)
	if err != nil {
		log.Fatal(err)
	}
	*/

}
