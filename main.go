package main

import (
	"log"
	"net/http"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/db"
	"github.com/shopcart/views"
)

func productFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("product func")

}

var app apiserver.AppI = &apiserver.App{}

func main() {
	// log.Println(models.DataModels.Products)
	db.InitDbApp()
	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("test here")
	})
	app.Get("/test", productFunc)
	views.InitAuthApp(app)
	views.InitStoreApp(app)
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
