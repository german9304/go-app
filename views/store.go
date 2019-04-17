package views

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
	"github.com/shopcart/models"
)

var productPath = "products"

var joinProductFiles = map[string]string{
	"product":       path.Join(productPath, "product.html"),
	"products":      path.Join(productPath, "products.html"),
	"topProducts":   path.Join(productPath, "topstores.html"),
	"createProduct": path.Join(productPath, "createproduct.html"),
}

var productTemplates = helper.GenerateTemplatePath("base.html", joinProductFiles)

// TODO: Finish product routes
//
func products(w http.ResponseWriter, r *http.Request) {
	pt := productTemplates["products"]
	productsModel := models.DataModels
	helper.RenderTemplate(w, pt, productsModel)
}

// createProduct creates a new Product with HTTP POST
func createProduct(w http.ResponseWriter, r *http.Request) {
	pt := productTemplates["createProduct"]
	productsModel := models.DataModels
	if r.Method == "POST" {
		fmt.Println("create product")
		// r.ParseForm()
		name := r.FormValue("name")
		quantity := r.FormValue("quantity")
		description := r.FormValue("description")

		qty, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal(err)
		}

		product := models.Product{Name: name, Quantity: qty, Description: description}
		fmt.Printf("New Product: %v", product)

		fmt.Printf("quantity is %v \n", quantity)
	}
	helper.RenderTemplate(w, pt, productsModel)
}

func topProducts(w http.ResponseWriter, r *http.Request) {
	pt := productTemplates["topProducts"]
	productsModel := models.DataModels
	helper.RenderTemplate(w, pt, productsModel)
}

// InitStoreApp initializes products app, adapter pattern
func InitStoreApp(app apiserver.AppI) {
	// p := path.Join("templates/app", "products")
	// fmt.Println(curr)
	app.Get("/products/", products)
	app.Get("/", products)
	app.Route("/create-product/", createProduct)
	app.Get("/top-products/", topProducts)
}
