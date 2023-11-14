package main

import (
	//"fmt"
	"html/template"
	"net/http"
	//"strings"
)


// Index Page handler
func index_page(w http.ResponseWriter, r *http.Request) {
	// Response template
	tmpl, _ := template.ParseFiles("www/index.html")
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}

