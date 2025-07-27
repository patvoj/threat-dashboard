package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

const (
	DefaultThreatsFile  = "internal/threats.jsonl"
	DefaultStaticDir    = "./ui/static"
	DefaultTemplatePath = "./ui/templates/threat.html.tmpl"
	DefaultPort         = ":4000"
)

type application struct {
	logger   *slog.Logger
	template *template.Template
	dataFile string
}

func main() {
	port := flag.String("p", DefaultPort, "port number")
	templ := flag.String("t", DefaultTemplatePath, "html template file path")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	template, err := template.ParseFiles(*templ)
	if err != nil {
		logger.Error("Failed to parse templates", "error", err, "template", *templ)
		os.Exit(1)
	}

	app := application{
		logger:   logger,
		template: template,
		dataFile: DefaultThreatsFile,
	}

	logger.Info("Starting a server at: ", "port", *port)

	if err := http.ListenAndServe(*port, app.routes()); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
