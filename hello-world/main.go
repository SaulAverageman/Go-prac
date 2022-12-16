package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who dares call home?")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "We do nothing.")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8000", nil)
}