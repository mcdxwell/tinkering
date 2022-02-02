package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TO-DO:

// Refer to this: https://projecteuler.net/thread=4;page=5

func isPalindrome(x int) bool {

	s := strconv.Itoa(x)
	fmt.Println("First s:", s)
	fmt.Println("Len(s): ", len(s))

	// append each character of s to slice
	// slice is currently empty.
	slice := make([]string, len(s))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Println(slice)
	ss := strings.Join(slice, ",")
	fmt.Println(ss)
	fmt.Println("S:", s)
	return s == ss
}

func main() {

	isPalindrome(10001)
}
