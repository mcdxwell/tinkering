package main

import (
	"fmt"
	"math"
)

func sumOfSquares() float64 {
	sum := 0.0
	for i := 0.0; i <= 100; i++ {
		sum += (math.Pow(i, 2))
	}
	return sum
}

func squareOfSum() float64 {
	sum := 0.0
	for i := 0.0; i <= 100; i++ {
		sum += i
	}
	return math.Pow(sum, 2)
}

func main() {
	diff := squareOfSum() - sumOfSquares()
	fmt.Println("Difference:", diff)
}
