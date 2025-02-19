package main

import (
	"log/slog"
	"net/http"
)

type httpErrors struct {
	logger *slog.Logger
}

func newHTTPErrors(logger *slog.Logger) httpErrors {
	return httpErrors{
		logger: logger,
	}
}

func (h httpErrors) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	h.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	h.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (h httpErrors) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{
		"error": message,
	}

	err := writeJSON(w, status, env, nil)
	if err != nil {
		h.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h httpErrors) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	h.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (h httpErrors) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	h.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (h httpErrors) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	h.errorResponse(w, r, http.StatusNotFound, message)
}

func (h httpErrors) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	h.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (h httpErrors) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	h.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (h httpErrors) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	h.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (h httpErrors) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	h.errorResponse(w, r, http.StatusForbidden, message)
}

func (h httpErrors) notPermittedResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	h.errorResponse(w, r, http.StatusForbidden, message)
}

func (h httpErrors) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	h.logger.Error(err.Error(), "method", method, "uri", uri)
}
