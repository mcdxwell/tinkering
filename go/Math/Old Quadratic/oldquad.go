package main

import (
	"fmt"
	"math"
)

func oldQuad(a float64, b float64, c float64) (float64, float64) {

	x1 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a
	x2 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / 2 * a

	return x1, x2
}

func main() {

	x1, x2 := oldQuad(1, 5, 6)
	fmt.Println(x1)
	fmt.Println(x2)
}
