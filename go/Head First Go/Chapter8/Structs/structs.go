package main

import "fmt"

func main() {

	var myStruct struct {
		number  float64
		word    string
		enabled bool
	}
	fmt.Printf("%#v\n", myStruct)
}
