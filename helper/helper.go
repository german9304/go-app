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

// GenerateTemplatePath generates templates for list of paths
func GenerateTemplatePath(baseTemplate string, files map[string]string) map[string]*template.Template {
	templates := make(map[string]*template.Template)
	for k, v := range files {
		templates[k] = template.Must(template.ParseFiles(baseTemplate, v))
	}
	return templates
}
