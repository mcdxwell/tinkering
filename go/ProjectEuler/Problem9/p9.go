package main

import "fmt"

// https://projecteuler.net/problem=9

func test_triplet(a int, b int, c int) bool {
	return (a*a + b*b) == c*c
}

func get_triplets(limit int) {
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

			if test_triplet(a, b, c) {
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
	get_triplets(1000)
}
