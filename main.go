package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world handler function
	hello_world := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", hello_world)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
