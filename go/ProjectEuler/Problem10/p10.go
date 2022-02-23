package main

import (
	"fmt"
)

func sieveOfAtkins(limit int) {
	if limit > 2 {
		fmt.Printf("2 ")
	}

	if limit > 3 {
		fmt.Printf("3 ")
	}

	sieve := make([]bool, limit)

	for i := 0; i < limit; i++ {
		sieve[i] = false
	}

	for x := 1; x*x < limit; x++ {
		for y := 1; y*y < limit; y++ {

			n := 4*x*x + y*y
			if n <= limit && n%12 == 1 || n%12 == 5 {

				// uncertain
				if sieve[n] != true {
					sieve[n] = true
				}
			}

			n = 3*x*x + y*y
			if n <= limit && n%12 == 7 {
				if sieve[n] != true {
					sieve[n] = true
				}
			}

			n = 3*x*x - y*y
			if x > y && n <= limit && n%12 == 11 {
				if sieve[n] != true {
					sieve[n] = true
				}
			}
		}
	}

	for r := 5; r*r < limit; r++ {
		if sieve[r] {
			for i := r * r; i < limit; i += r * r {
				sieve[i] = false
			}
		}
	}

	for a := 5; a < limit; a++ {
		if sieve[a] {
			fmt.Printf("%d ", a)
		}
	}

}

func main() {
	sieveOfAtkins(2000000)
}
