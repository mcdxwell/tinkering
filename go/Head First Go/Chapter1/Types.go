package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(reflect.TypeOf(69))
	fmt.Println(reflect.TypeOf(420.69))
	fmt.Println(reflect.TypeOf("420.69"))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('A'))

	// Ouput: int
	// Ouput: float64
	// Ouput: string
	// Ouput: string
	// Ouput: bool
	// Ouput: int32

}
