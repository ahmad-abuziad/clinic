package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Listening to :4000")

	err := http.ListenAndServe(":4000", routes())

	fmt.Println(err.Error())
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/patient", createPatientHandler)

	return mux
}
