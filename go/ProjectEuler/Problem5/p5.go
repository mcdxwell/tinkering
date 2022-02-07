package main

import (
	"fmt"
	"math"
)

func sieve() int {

	k := 20
	n := 0
	i := 1

	check := true

	limit := math.Sqrt(float64(k))

	p := [8]int{2, 3, 5, 7, 11, 13, 17, 19}
	a := make([]int, 4)
	//var a []int
	for p[i] <= k {
		if check {
			if p[i] <= int(limit) {
				a[i] = int(math.Floor(math.Log(float64(k)) / math.Log(float64(p[i]))))
				fmt.Println(a[i])
			} else {
				check = false
			}
		}

		n *= int(math.Pow(float64(p[i]), float64(a[i])))
		i += 1
	}
	fmt.Println(n)
	return n
}

func main() {
	sieve()
}
