package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.text")

	if fileInfo != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.Size())
}
