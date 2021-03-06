package main

import "fmt"

func testTriplet(a, b, c int) bool {
	return (a*a + b*b) == c*c
}

func getTriplets(limit int) {
	a, b, c := 0, 0, 0
	m := 2

	for c < limit {
		for n := 1; n <= m-1; n++ {

			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n

			if c > limit {
				break
			}

			if testTriplet(a, b, c) {
				if a+b+c == 1000 {
					prod := a * b * c
					fmt.Printf("The product of %d %d %d = %d", a, b, c, prod)
				}
			}
			//fmt.Printf("%d %d %d\n", a, b, c)
		}
		m++
	}
}

func main() {
	getTriplets(1000)
}
