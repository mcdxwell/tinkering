package main

import (
	"fmt"
	"math"
	"strconv"
)

// TO-DO:

// Refer to this: https://projecteuler.net/thread=4;page=5
// reference: https://go.dev/play/p/5FUOzjBa-o
// https://stackoverflow.com/questions/52449646/array-with-integers-and-nil-values-in-golang/52450575
func isPalindrome(input int) bool {
	x := strconv.Itoa(input)
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-i-1] {
			return false
		}
	}

	return true
}

type NilInt struct {
	value int
	null  bool
}

func (n *NilInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(x int) NilInt {
	return NilInt{x, false}
}

func NewNil() NilInt {
	return NilInt{0, true}
}

func inverse(x, mod int) NilInt {
	/* one := 1
	sev := 7
	thr := 3
	nin := 9 */

	//ints := []*int{nil, &one, nil, &sev, nil, nil, nil, &thr, nil, &nin}
	//remainder := ints[x%10]
	var y = []NilInt{NewNil(), NewInt(1), NewNil(), NewInt(7), NewNil(), NewNil(), NewNil(), NewInt(3), NewNil(), NewInt(9)}
	fmt.Println(y)
	remainder := y[x%10]
	fmt.Println("First remainder: \n", remainder)
	/* for _, i := range y {
		fmt.Printf("%v ", i.Value())
	} */
	/* for _, anInt := range ints {
		if anInt != nil {
			fmt.Println("Rem1:", remainder)
			fmt.Println("*anInt:", *anInt)
		} else {
			fmt.Println("Rem2:", remainder)
			fmt.Println("anInt:", anInt)
		}
	} */

	if remainder.value == 0 {
		return remainder
	}

	for {
		ax := remainder.value * x % mod

		if ax == 1 {
			return remainder
		}

		remainder.value = (remainder.value * (2 - (ax)))
		fmt.Println("Rem val lol:remainder.value \n", remainder.value)
	}

}

func pas2(n float64) {

	if n < 2 {
		panic("n is less than 2")
	}
	k := n / 2
	for {
		maxf := math.Pow(10, n) - 1
		maxf11 := (maxf-11)/22*22 + 11
		minf := math.Pow(10, n) - math.Pow(10, n-k) + 1

		if 2*k == n {
			best := maxf * minf
		}

	}
}

func main() {
	inverse(2000, 10)
	isPalindrome(10001)

}
