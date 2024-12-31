package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8081"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Printf("Listening on port %v\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
