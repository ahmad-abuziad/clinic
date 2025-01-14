package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	port   string
	logger *slog.Logger
}

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		port:   ":4000",
		logger: logger,
	}

	logger.Info("Starting server")

	err := http.ListenAndServe(app.port, app.routes())

	fmt.Println(err.Error())
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patient", app.createPatientHandler)

	return mux
}
