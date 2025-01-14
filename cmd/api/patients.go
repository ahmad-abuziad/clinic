package main

import (
	"net/http"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/data"
)

func (app *application) createPatientHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		DateOfBirth time.Time `json:"date_of_birth"`
		Gender      string    `json:"gender"`
		Notes       string    `json:"notes"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	patient := &data.Patient{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		DateOfBirth: input.DateOfBirth,
		Gender:      input.Gender,
		Notes:       input.Notes,
	}

	if errors := data.ValidatePatient(patient); len(errors) > 0 {
		app.failedValidationResponse(w, r, errors)
		return
	}

	writeJSON(w, http.StatusCreated, envelope{"patient": patient}, nil)
}
