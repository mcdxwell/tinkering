package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a := getRandNums(8)
	fmt.Println("Unsorted: \n", a)
	b := quicksort(a)
	fmt.Println("Sorted: \n", b)

}

func getRandNums(size int) []int {
	// using the data structure slice
	a := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		a[i] = rand.Intn(100) - rand.Intn(100)
	}
	return a
}

// refer to Chapter 7.1 - Introduction to Algorithms 3rd ed.
func quicksort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	l, r := 0, len(a)-1

	p := rand.Int() % len(a)

	a[p], a[r] = a[r], a[p]

	for i := range a {
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l]
			l++
		}
	}

	a[l], a[r] = a[r], a[l]
	quicksort(a[:l])
	quicksort(a[l+1:])

	return a
}
