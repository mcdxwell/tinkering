package main

import "fmt"

// we can use pointers to allow a function to update a struct

type sub struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s *sub) {
	s.rate = 4.99
}

func main() {
	var s sub
	applyDiscount(&s)
	fmt.Println(s.rate)
}
