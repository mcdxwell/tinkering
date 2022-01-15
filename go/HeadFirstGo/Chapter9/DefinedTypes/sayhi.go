package main

import "fmt"

type MyType string

// m is the receiver parameter
// MyType is the receiver parameter type
func (m MyType) sayHi() {
	// m - > print the receiver parameter's value
	fmt.Println("Hi, from", m)
}

func main() {
	value := MyType("David")

	// Receivers passed to receiver parameter
	value.sayHi()

	anotherValue := MyType("Hank Hill")

	// Receivers passed to receiver parameter
	anotherValue.sayHi()
}

/* Output:
// Receiver values are in output (David) (Hank Hill)
Hi, from David
Hi, from Hank Hill
*/

// Think about receiver parameters like *this* and *self* values in other languages