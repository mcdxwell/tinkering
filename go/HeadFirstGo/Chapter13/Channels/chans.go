package main

import "fmt"

func foo(channel chan string) {
	channel <- "food"
	channel <- "rood"
	channel <- "dood"
}

func bar(channel chan string) {
	channel <- "rawr"
	channel <- "czar"
	channel <- "car"
}

func baz(channel chan string) {
	channel <- "taz"
	channel <- "ras"
	channel <- "das"
}

func main() {
	oofChan := make(chan string)
	rabChan := make(chan string)
	zabChan := make(chan string)

	go foo(rabChan)
	go bar(zabChan)
	go baz(oofChan)

	fmt.Println(<-oofChan)
	fmt.Println(<-oofChan)
	fmt.Println(<-oofChan)

	fmt.Println(<-zabChan)
	fmt.Println(<-zabChan)
	fmt.Println(<-zabChan)

	fmt.Println(<-rabChan)
	fmt.Println(<-rabChan)
	fmt.Println(<-rabChan)
	fmt.Println()
}
