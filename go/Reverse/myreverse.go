package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := []int{5, 2, 6, 3, 1, 4, 69, 10, 11, 12, 1263, 15645, 1231324564, 564, 123123, 456, 123412485, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 6}
	sLen := len(s)
	roons := make([]int, 0)
	for i := sLen - 1; i >= 0; i-- {
		roons = append(roons, s[i])
	}

	fmt.Println(roons)
	duration := time.Since(start)
	fmt.Println(duration.Milliseconds())
}
