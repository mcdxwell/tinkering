package main

import "fmt"

func main() {

	// declaring the variables
	var quantity int
	var length, width float64
	var customerName string

	// assigning the values to the variables
	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
