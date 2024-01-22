package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error passing html file err : %v", err))
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("error exucting html  err : %v", err))

	}
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", "")
}
