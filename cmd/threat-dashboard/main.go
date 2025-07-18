package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

const (
	DefaultThreatsFile = "internal/threats.jsonl"
	DefaultStaticDir   = "./ui/static"
)

type application struct {
	logger    *slog.Logger
	templates *template.Template
	dataFile  string
}

func main() {
	port := flag.String("p", ":4000", "port number")
	templ := flag.String("t", "./ui/templates/threat.html.tmpl", "html template file path")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	templates, err := template.ParseFiles(*templ)
	if err != nil {
		logger.Error("Failed to parse templates", "error", err, "template", *templ)
		os.Exit(1)
	}

	app := application{
		logger:    logger,
		templates: templates,
		dataFile:  DefaultThreatsFile,
	}

	logger.Info("Starting a server at: ", "port", *port)

	if err := http.ListenAndServe(*port, app.routes()); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
