package main

import "fmt"

// we can use pointers to allow a function to update a struct
// pass large structs using pointers
type sub struct {
	name   string
	rate   float64
	active bool
}

// this function uses *sub to take a pointer
func printInfo(s *sub) {
	fmt.Println("Name: ", s.name)
	fmt.Println("Rate: ", s.rate)
	fmt.Println("Active: ", s.active)
}

func defaultSub(name string) *sub { // *sub returns pointer
	var s sub
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // return a pointer to a struct instead of the struct itself
}

func applyDiscount(s *sub) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSub("David McDowell")
	applyDiscount(subscriber1)
	printInfo(subscriber1)

	subscriber2 := defaultSub("Hank Hill")
	printInfo(subscriber2)
}

// struct names and struct fields must be capitalized if you want to export them
// from their package
