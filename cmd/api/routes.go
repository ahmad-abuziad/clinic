package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patients", app.requirePermission("patient:write", app.createPatientHandler))
	mux.HandleFunc("GET /v1/patients/{id}", app.requirePermission("patient:read", app.getPatientHandler))
	mux.HandleFunc("POST /v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.authenticate(mux)
}
