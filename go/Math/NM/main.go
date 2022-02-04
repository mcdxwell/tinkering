package main

import (
	"fmt"
	"math"
)

func fx(x float64) float64 {
	//fx := math.Sin(math.Pow(x, 3)) + math.Pow(x, 2)
	fx := math.Pow(x, 2) - math.Sin(2*math.Pow(x, 2)) - 1
	return fx
}

func fxh(x, h float64) float64 {
	//fxh := math.Sin(math.Pow(x+h, 3)) + math.Pow(x+h, 2)
	fxh := math.Pow(x+h, 2) - math.Sin(2*math.Pow(x+h, 2)) - 1
	return fxh
}

func funky() {
	h := 2.0
	//x := 1.0
	x := 2.0
	bound := 2.0e-12

	//trueAnsw := 2 + 3*math.Cos(1)
	trueAnsw := 4 - 8*math.Cos(8)
	for i := 0; h > bound; i++ {
		h /= 2
		pp := (fxh(x, h) - fx(x)) / h
		errEst := trueAnsw - pp
		fmt.Printf("%v) h: %v F(x)': %v Error: %v\n", i, h, pp, errEst)
	}
}

func main() {
	funky()
}
