package views

import (
	"net/http"
	"os"
	"path"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
	"github.com/shopcart/models"
)

var currAuth, _ = os.Getwd()
var baseAuthTemplate = path.Join(currAuth, "templates", "base.html")

var authPath = path.Join(currAuth, "templates", "auth")

var joinAuthFiles = map[string]string{
	"login":    path.Join(authPath, "login.html"),
	"register": path.Join(authPath, "register.html"),
}

var authTemplates = helper.GenerateTemplatePath(baseAuthTemplate, joinAuthFiles)

// login handler
func login(w http.ResponseWriter, r *http.Request) {
	lt := authTemplates["login"]
	helper.RenderTemplate(w, lt, models.DataModels)
}

// register hanlder
func register(w http.ResponseWriter, r *http.Request) {
	rt := authTemplates["register"]
	helper.RenderTemplate(w, rt, models.DataModels)

}

// InitAuthApp exports init auth app pattern
func InitAuthApp(app apiserver.AppI) {
	app.Get("/login/", login)
	app.Get("/register/", register)
}
