package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello, World!")
		fmt.Fprint(rw, "Learning WSL2 😀\n")
	})

	http.ListenAndServe(":8080", nil)
}
