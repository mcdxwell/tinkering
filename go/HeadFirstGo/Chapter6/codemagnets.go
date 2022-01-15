package main

import "fmt"

func sum(numbers ...int) int {

	var sum int = 0

	for _, number := range numbers {

		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sum(7, 9))
	fmt.Println(sum(1, 2, 4))
}
