package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/data"
)

func main() {

	fmt.Println("Listening to :4000")

	err := http.ListenAndServe(":4000", routes())

	fmt.Println(err.Error())
}

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	patient := data.Patient{
		FirstName:   "Ahmad",
		LastName:    "Abuziad",
		DateOfBirth: time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC),
		Gender:      "M",
		Notes:       "Gluten Allergic",
	}

	js, _ := json.Marshal(envelope{"patient": patient})

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patient", createPatientHandler)

	return mux
}
