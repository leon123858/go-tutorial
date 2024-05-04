package main

import "net/http"

func main() {
	// print on localhost:8080
	println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("kitchen")))
	if err != nil {
		return
	}
}
