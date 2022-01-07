package main

import (
	"fmt"
	"math"
)

func newQuad(b float64, c float64) (float64, float64) {

	quadPos := -b/2 + math.Sqrt(((math.Pow(b, 2))/4)-c)
	quadNeg := -b/2 - math.Sqrt(((math.Pow(b, 2))/4)-c)
	return quadPos, quadNeg
}

func main() {

	a, b := newQuad(2, -15)
	fmt.Println(a)
	fmt.Println(b)
}
