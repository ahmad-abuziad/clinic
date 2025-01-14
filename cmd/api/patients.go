package main

import (
	"net/http"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/data"
)

func createPatientHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		DateOfBirth time.Time `json:"date_of_birth"`
		Gender      string    `json:"gender"`
		Notes       string    `json:"notes"`
	}

	err := readJSON(w, r, &input)
	if err != nil {
		badRequest(w, err)
		return
	}

	// validate

	patient := data.Patient{
		FirstName:   "Ahmad",
		LastName:    "Abuziad",
		DateOfBirth: time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC),
		Gender:      "M",
		Notes:       "Gluten Allergic",
	}

	writeJSON(w, http.StatusCreated, envelope{"patient": patient}, nil)
}
