package helper

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate is a helper function to render a template
func RenderTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {

	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
