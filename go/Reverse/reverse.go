package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := []int{5, 2, 6, 3, 1, 4, 69, 10, 11, 12, 1263, 15645, 1231324564, 564, 123123, 456, 123412485, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)
	duration := time.Since(start)
	fmt.Println(duration.Milliseconds())
}
