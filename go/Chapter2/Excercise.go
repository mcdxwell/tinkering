package main

import (
	"fmt"
)

// comment out the lines that have errors
// problem: some lines refer to a variable that is out of scope
var a = "a"

func main() {
	a = "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println(a)
			fmt.Println(b)
			fmt.Println(c)
			fmt.Println(d)
		}
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		//fmt.Println(d)
	}
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
}
