package apiserver

import (
	"log"
	"net/http"
)

// AppI interface for App
type AppI interface {
	Get(pattern string, fn http.HandlerFunc)
	Post(pattern string, fn http.HandlerFunc)
	Route(pattern string, fn http.HandlerFunc)
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

func (app *App) makeRouteHanlder(fn http.HandlerFunc) http.HandlerFunc {
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

// Route method from app
func (app *App) Route(pattern string, fn http.HandlerFunc) {
	http.HandleFunc(pattern, app.makeRouteHanlder(fn))
}
