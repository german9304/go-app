package main

import (
	"log"
	"net/http"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/views"
)

func productFunc(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// log.Println(models.DataModels.Products)
	var app apiserver.AppI = &apiserver.App{}
	app.Get("/test", productFunc)
	views.InitStoreApp(app)
	views.InitAuthApp(app)
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
