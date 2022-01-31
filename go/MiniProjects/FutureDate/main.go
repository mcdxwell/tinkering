/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"FutureDate/cmd"
	"fmt"
)

func main() {
	fmt.Println("The current date is: ", cmd.GetCurrDate().Format("01-02-2006"))
	cmd.Execute()
}
