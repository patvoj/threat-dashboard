package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(DefaultStaticDir))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", app.index)
	mux.HandleFunc("/render", app.render)

	app.logger.Debug("Routes configured successfully")
	return mux
}
