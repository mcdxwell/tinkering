package main

import (
	"fmt"
	"math"
)

// This program is based on the overview pseudocode provided by ProjectEuler

func isPrime(n int) bool {
	if n == 1 {
		return false
	} else if n < 4 {
		return true
	} else if n%2 == 0 {
		return false
	} else if n < 9 {
		return true
	} else if n%3 == 0 {
		return false
	} else {
		r := math.Floor(math.Sqrt(float64(n)))
		f := 5.0

		for f <= r {
			if n%int(f) == 0 {
				return false
			}

			if n%(int(f)+2) == 0 {
				return false
			}
			f += 6
		}
		return true
	}
}

func funky() int {
	limit := 10001
	count := 1
	candidate := 1

	for count < limit {
		candidate += 2
		if isPrime(candidate) {
			count += 1
		}
	}
	fmt.Println(candidate)
	return candidate
}
func main() {
	funky()
}
