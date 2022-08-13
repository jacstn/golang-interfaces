package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".tmpl", "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "This is about page")
	}
}
