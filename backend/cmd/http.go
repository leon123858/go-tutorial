package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		return
	}
}
