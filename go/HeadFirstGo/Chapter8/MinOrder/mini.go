package main

import "fmt"

type part struct {
	description string
	count       int
}

// Declare one return value (type -> part)
func minimumOrder(description string) part {
	var p part // create a new "part" value
	p.description = description
	p.count = 100
	return p // return the part
}

func main() {
	p := minimumOrder("Hex bolts") // call minimumOrder ()
	fmt.Println(p.description, p.count)
}
