package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := flag.String("p", ":4000", "port number")
	// templ := flag.String("t", "", "html template file path")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	app := application{
		logger: logger,
	}

	logger.Info("Starting a server at port: ", *port)

	err := http.ListenAndServe(*port, app.routes())
	logger.Error("Server could not start", err.Error())
	os.Exit(1)
}

type application struct {
	logger *slog.Logger
}
