package main

import "fmt"

// The sum of all even fibonacci numbers less than 4 Million
func evenFib(a int, b int) int {

	sum := 0
	for a < 4000000 {
		if a%2 == 0 {
			sum += a
		}
		a, b = b, a // swap, works just like python
		a += b
	}
	return sum
}

func main() {
	fmt.Printf("Sum: %d", evenFib(1, 1))
}
