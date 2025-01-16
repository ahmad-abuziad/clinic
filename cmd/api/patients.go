package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/data"
	"github.com/ahmad-abuziad/clinic/internal/validator"
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

	v := validator.New()

	if data.ValidatePatient(v, patient); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Patients.Insert(patient)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/patients/%d", patient.ID))

	err = writeJSON(w, http.StatusCreated, envelope{"patient": patient}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPatientHandler(w http.ResponseWriter, r *http.Request) {
	id, err := readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	patient, err := app.models.Patients.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"patient": patient}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
