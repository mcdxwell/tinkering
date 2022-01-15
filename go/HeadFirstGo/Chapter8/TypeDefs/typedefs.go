package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	fmt.Println("Name: ", porsche.name)
	fmt.Println("Top speed: ", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description: ", bolts.description)
	fmt.Println("Count: ", bolts.count)
}
