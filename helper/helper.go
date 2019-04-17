package helper

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	curr         string
	err          error
	baseTemplate string
)

// RenderTemplate is a helper function to render a template
func RenderTemplate(w http.ResponseWriter, t *template.Template, data interface{}) {

	err := t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateTemplatePath generates templates for list of paths
func GenerateTemplatePath(baseT string, files map[string]string) map[string]*template.Template {
	bt := baseTemplate
	joinBasePaths := path.Join(bt, baseT)
	templates := make(map[string]*template.Template)
	for k, v := range files {
		joinTemplatePaths := path.Join(bt, v)
		templates[k] = template.Must(template.ParseFiles(joinBasePaths, joinTemplatePaths))
	}
	return templates
}

func init() {
	curr, err = os.Getwd()                      // get current directory
	baseTemplate = path.Join(curr, "templates") //base template folder
}
