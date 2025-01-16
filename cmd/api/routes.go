package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patients", app.createPatientHandler)
	mux.HandleFunc("GET /v1/patients/{id}", app.getPatientHandler)

	return mux
}
