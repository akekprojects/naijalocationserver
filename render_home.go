package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, data interface{}) {
	htmlString := `
	<html>
    <head>
        <h1>
            Hi this is Naija Location server endpoint.
          
        </h1>
        
    </head>


    <p >
        To read about active endpoints and how to get started check github README <a href="https://github.com/akekprojects/naijalocationserver/blob/main/README.md"> here</a>
    </p>

	<p>
	Note: This site is still under construction and is being worked on.
	test
</p>
</html>
	`
	t, err := template.New("index").Parse(htmlString)
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
	renderTemplate(w, "")
}
