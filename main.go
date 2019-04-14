package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
	"github.com/shopcart/models"
)

var pT = template.Must(template.ParseFiles("./templates/base.html", "./templates/products/products.html"))

func productFunc(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, pT, models.DataModels)

}

func main() {
	log.Println(models.DataModels.Products)
	var app apiserver.AppI = &apiserver.App{}
	app.Get("/", productFunc)
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
