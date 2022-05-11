package main

import "fmt"

// Lattice Paths
// https://en.wikipedia.org/wiki/Lattice_path

func lattice(n int) int {
	routes := 1
	for i := 1; i <= n; i++ {
		routes = routes * (n + i) / i
	}
	return routes
}

func main() {
	n := 20
	r := lattice(n)
	fmt.Printf("There would be %d routes on a %d by %d grid.", r, n, n)
}
