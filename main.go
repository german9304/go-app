package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.html")
	if err != nil {
		log.Fatal(err)
	}
	helper.RenderTemplate(w, t, "")
}

func aboutFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "about %v ", "page")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var app apiserver.AppI = &apiserver.App{}
	app.Get("/", helloFunc)
	app.Get("/about/", aboutFunc)
	err := http.ListenAndServe(":8081", nil)
	log.Fatal(err)
}
