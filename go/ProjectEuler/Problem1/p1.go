package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of all the multiples of (3 or 5) < 1000\n%d", sum)
}
