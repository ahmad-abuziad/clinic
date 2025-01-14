package main

import "net/http"

func badRequest(w http.ResponseWriter, err error) {
	env := envelope{
		"error": err.Error(),
	}

	writeJSON(w, http.StatusBadRequest, env, nil)
}
