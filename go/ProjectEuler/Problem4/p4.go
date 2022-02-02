package main

import (
	"strconv"
)

// TO-DO:

// Refer to this: https://projecteuler.net/thread=4;page=5
// reference: https://go.dev/play/p/5FUOzjBa-o
func isPalindrome(input int) bool {
	x := strconv.Itoa(input)
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-i-1] {
			return false
		}
	}

	return true
}

func main() {

	isPalindrome(10001)
}
