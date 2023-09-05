package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.HandleFunc("/", app.home)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/jerry/create-snippet", app.snippetCreate)
	mux.HandleFunc("/jerry/view-snippet", app.snippetView)
	return mux
}
