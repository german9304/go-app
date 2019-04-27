package views

import (
	"database/sql"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/shopcart/apiserver"
	"github.com/shopcart/helper"
	"github.com/shopcart/models"
)

// var currAuth, _ = os.Getwd()
// var baseAuthTemplate = path.Join(currAuth, "templates", "base.html")

// TODO: Implement logout route

var authPath = "auth"

var joinAuthFiles = map[string]string{
	"login":    path.Join(authPath, "login.html"),
	"register": path.Join(authPath, "register.html"),
}

var authTemplates = helper.GenerateTemplatePath("base.html", joinAuthFiles)

// login handler
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		log.Println(email, password)
		db, ok := apiserver.Global["db"]
		if !ok {
			log.Fatal("element not found")
		}
		user, err := models.GetUser(db.(*sql.DB), email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user)
		// Setting the cookie for session and testing
		if user.Password == password {
			userID := int(user.ID)
			cookie := http.Cookie{Name: "userID", Value: strconv.Itoa(userID)}
			log.Println(cookie.Value)
			log.Println("user is authenticated")
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/products", http.StatusFound)
		}
	}
	lt := authTemplates["login"]
	helper.RenderTemplate(w, lt, models.DataModels)
}

// register hanlder
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		log.Printf("Username: %v, email: %v, password: %v \n", username, email, password)
		log.Println("redirecting")
		db, ok := apiserver.Global["db"]
		if !ok {
			log.Fatal("element not found")
		}
		row, user, err := models.CreateUser(db.(*sql.DB), email, username, password)
		if err == nil {
			log.Println(row, user)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		log.Println("user already exits", err)
	}
	rt := authTemplates["register"]
	helper.RenderTemplate(w, rt, models.DataModels)

}

func logout(w http.ResponseWriter, r *http.Request) {
	// logout
	cookie := http.Cookie{Name: "userID", MaxAge: -1}
	log.Println(cookie)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

// LoginRequired checks if user is athenticated middleware
func loginRequired(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			log.Println("login required")
			cookie, err := r.Cookie("userID")
			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}
			log.Println(cookie.Value)
		}
		fn(w, r)
	}
}

// InitAuthApp exports init auth app pattern
func InitAuthApp(app apiserver.AppI) {
	app.Get("/login", login)
	app.Get("/register", register)
	app.Get("/logout", logout)
}
