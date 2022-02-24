package main

import (
	"fmt"
	"math"
)

func sieveOfAtkins(ch chan int, limit int) {

	sieve := make([]bool, limit)
	limitSqrt := int(math.Sqrt(float64(limit)))
	for i := 0; i < limit; i++ {
		sieve[i] = false
	}

	var n int
	for x := 1; x <= limitSqrt; x++ {
		for y := 1; y <= limitSqrt; y++ {
			n = (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				sieve[n] = !sieve[n]
			}

			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				sieve[n] = !sieve[n]
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				sieve[n] = !sieve[n]
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

	ch <- 2
	ch <- 3
	for a := 5; a < limit; a++ {
		if sieve[a] {
			ch <- a
		}
	}
	close(ch)
}

func main() {
	prime_sum := uint64(0)
	ch := make(chan int)
	go sieveOfAtkins(ch, 2000000)
	for prime := range ch {
		prime_sum += uint64(prime)
	}
	fmt.Println(prime_sum)
}
