package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.index(app.templPath))
	mux.HandleFunc("POST /render", app.render)

	return mux
}
