package views

import (
	"net/http"
	"os"

	"path"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
	"github.com/shopcart/models"
)

var curr, _ = os.Getwd()
var baseTemplate = path.Join(curr, "templates", "base.html")

var productPath = path.Join(curr, "templates", "products")

var joinFiles = map[string]string{
	"product":       path.Join(productPath, "product.html"),
	"products":      path.Join(productPath, "products.html"),
	"topProducts":   path.Join(productPath, "topstores.html"),
	"createProduct": path.Join(productPath, "createproduct.html"),
}

var productTemplates = helper.GenerateTemplatePath(baseTemplate, joinFiles)

//
func products(w http.ResponseWriter, r *http.Request) {
	var pt = productTemplates["products"]
	productsModel := models.DataModels
	helper.RenderTemplate(w, pt, productsModel)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var pt = productTemplates["createProduct"]
	productsModel := models.DataModels
	helper.RenderTemplate(w, pt, productsModel)
}

func topProducts(w http.ResponseWriter, r *http.Request) {
	var pt = productTemplates["topProducts"]
	productsModel := models.DataModels
	helper.RenderTemplate(w, pt, productsModel)
}

// InitApp initializes products app, adapter pattern
func InitApp(app apiserver.AppI) {
	app.Get("/products/", products)
	app.Get("/", products)
	app.Route("/create-product/", createProduct)
	app.Get("/top-products/", topProducts)
}
