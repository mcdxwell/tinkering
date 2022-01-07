package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(now)
	fmt.Println(year)

	broken := "G# R#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
