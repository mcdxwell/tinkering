package main

import "fmt"

// given 600851475143, what is its largest prime factor?
func main() {
	number := 600851475143
	j := 2
	for i := 0; i < number; i++ {
		if number%j == 0 {
			number /= j
		} else {
			j += 1
		}
	}
	fmt.Println(j)
}
