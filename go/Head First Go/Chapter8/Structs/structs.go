package main

import "fmt"

func main() {

	var myStruct struct {
		number  float64
		word    string
		enabled bool
	}
	fmt.Printf("%#v\n", myStruct)

	var subscriber struct {
		name   string
		rate   float64
		active bool
	}

	subscriber.name = "David"
	subscriber.rate = 2.81
	subscriber.active = true

	fmt.Println("Name: ", subscriber.name)
	fmt.Println("Monthly rate: ", subscriber.rate)
	fmt.Println("Active?: ", subscriber.active)
}
