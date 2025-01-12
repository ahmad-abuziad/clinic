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
	w.WriteHeader(http.StatusCreated)
	patient := data.Patient{
		FirstName:   "Ahmad",
		LastName:    "Abuziad",
		DateOfBirth: time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC),
		Gender:      "M",
		Notes:       "Gluten Allergic",
	}

	js, _ := json.Marshal(envelope{"patient": patient})

	w.Write(js)
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patient", createPatientHandler)

	return mux
}
