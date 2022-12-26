package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, path string) {
	page, _ := template.ParseFiles("./templates/" + path)
	err := page.Execute(w, nil)
	if err != nil {
		fmt.Println("error:", err)
	}
}
