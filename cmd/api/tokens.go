package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/data"
	"github.com/ahmad-abuziad/clinic/internal/validator"
)

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		app.errors.badRequest(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.Valid() {
		app.errors.failedValidationResponse(w, r, v.Errors)
		return
	}

	// get user by email

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.errors.invalidCredentialsResponse(w, r)
		default:
			app.errors.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.errors.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.errors.invalidCredentialsResponse(w, r)
	}

	token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		app.errors.serverErrorResponse(w, r, err)
	}

	err = writeJSON(w, http.StatusCreated, envelope{"token": token}, nil)
	if err != nil {
		app.errors.serverErrorResponse(w, r, err)
	}
}
