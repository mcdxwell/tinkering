package main

import "fmt"

// Long Collatz Sequence
// You have two operations and we need to do them recursively
// When n is even -> n/2
// When n is odd -> 3n+1
// This is performed until n equals 1

func collatz(n int) int {
	seq := 1
	for 1 < n {
		if n&1 != 1 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		seq += 1
	}
	return seq
}

func main() {

	chain, number := 0, 0
	for i := 2; i < 1000000; i++ {
		length := collatz(i)
		if length > chain {
			chain = length
			number = i
		}
	}
	fmt.Printf("Number: %d gave a chain of length: %d", number, chain)
}
