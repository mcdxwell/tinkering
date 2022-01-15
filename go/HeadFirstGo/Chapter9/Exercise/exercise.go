package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number) Double() {
	*n *= 2
}

func main() {
	ten := Number(10)
	ten.Add(4)
	ten.Subtract(4)

	four := Number(4)
	four.Add(3)
	four.Subtract(2)

	number := Number(4)
	fmt.Println("Original value of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)
}
