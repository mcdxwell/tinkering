// David 'Davz' McDowell
// 1-11-22
package main

import (
	"fmt"
	"time"
)

// To-DO:
// Give SetDay() parameters for time (date, hours, secs.)
func SetDay() []string {
	sats := []string{}
	// 5-4-22 -> First day of "SKILSTAK Beginner Boost"
	// twitch.tv/rwxrob
	t := time.Date(2022, 5, 4, 11, 11, 11, 11, time.UTC)
	for i := 1; i <= 16; i++ {
		ft := t.Format("Jan-02-2006 3:4 PM")
		sats = append(sats, ft)
		fmt.Printf("Week %2d | %v\n", i, ft)
		t = t.AddDate(0, 0, 7)
	}
	return sats
}

func main() {
	SetDay()
}
