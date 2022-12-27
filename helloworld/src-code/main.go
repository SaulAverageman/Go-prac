package main

import (
	"net/http"
)

func main() {
	const appPort = ":8000"
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	_ = http.ListenAndServe(appPort, nil)

}
