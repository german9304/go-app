package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shopcart/apiserver"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello %v ", "worlds")
	if err != nil {
		log.Fatal(err)
	}
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
