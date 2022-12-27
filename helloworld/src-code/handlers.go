package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html")
}
