package apiserver

import (
	"log"
	"net/http"
)

// appI interface for App
type appI interface {
	get(http.HandlerFunc)
	post(http.HandlerFunc)
}

// App structure reprensets http methods
type App struct {
}

func (app *App) makeGetHanlder(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		fn(w, r)
	}
}

func (app *App) makePostHanlder(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		fn(w, r)
	}
}

// Get method from app
func (app *App) Get(pattern string, fn http.HandlerFunc) {
	http.HandleFunc(pattern, app.makeGetHanlder(fn))
}

// Post method from app
func (app *App) Post(pattern string, fn http.HandlerFunc) {
	http.HandleFunc(pattern, app.makePostHanlder(fn))
}
