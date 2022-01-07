package main

import (
	"fmt"
	"math"
)

func newQuad(b float64, c float64) (float64, float64) {

	x1 := -b/2 + math.Sqrt(((math.Pow(b, 2))/4)-c)
	x2 := -b/2 - math.Sqrt(((math.Pow(b, 2))/4)-c)
	return x1, x2
}

func main() {

	x1, x2 := newQuad(5, 6)
	fmt.Println(x1)
	fmt.Println(x2)
}
