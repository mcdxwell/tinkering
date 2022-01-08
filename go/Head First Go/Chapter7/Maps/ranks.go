package main

import "fmt"

func main() {
	/* var ranks map[string]int
	ranks = make(map[string]int)
	*/
	ranks := make(map[string]int) // create a map AND declare a variable to hold it
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])
	fmt.Println(ranks["silver"])
}
