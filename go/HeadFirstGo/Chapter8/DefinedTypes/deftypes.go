package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description: ", p.description)
	fmt.Println("Count: ", p.count)
}
func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)

	var gears part
	gears.description = "Hex gears"
	gears.count = 26
	showInfo(gears)
}
