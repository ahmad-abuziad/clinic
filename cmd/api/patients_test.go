package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/assert"
	"github.com/ahmad-abuziad/clinic/internal/data"
)

func TestCreatePatientHandler(t *testing.T) {
	app, _ := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	urlPath := "/v1/patients"

	t.Run("201 Created", func(t *testing.T) {
		var (
			validFirstName   = "Ahmad"
			validLastName    = "Abuziad"
			validDateOfBirth = time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC)
			validGender      = "M"
		)

		reqBody := fmt.Sprintf(`
		{
			"first_name": %q,
			"last_name": %q,
			"date_of_birth": %q,
			"Gender": %q
		}`, validFirstName, validLastName, validDateOfBirth.Format(time.RFC3339), validGender)

		statusCode, headers, body := ts.postJSON(t, urlPath, reqBody)

		assert.Equal(t, statusCode, http.StatusCreated)
		assert.Equal(t, headers.Get("Location"), "/v1/patients/2")

		gotPatient := unmarshalPatient(t, body)
		assert.Equal(t, gotPatient.FirstName, validFirstName)
		assert.Equal(t, gotPatient.LastName, validLastName)
		assert.Equal(t, gotPatient.Gender, validGender)
		assert.Equal(t, gotPatient.DateOfBirth, validDateOfBirth)
		assert.Equal(t, gotPatient.ID, 2)
		assert.Equal(t, gotPatient.CreatedAt.IsZero(), true)
	})

	t.Run("400 Bad Request", func(t *testing.T) {
		statusCode, _, body := ts.postJSON(t, urlPath, "{'invalid': json")

		assert.Equal(t, statusCode, http.StatusBadRequest)
		assert.StringContains(t, string(body), "error")
	})

	t.Run("422 Unprocessable Content", func(t *testing.T) {
		statusCode, _, body := ts.postJSON(t, urlPath, `{"first_name":""}`)

		assert.Equal(t, statusCode, http.StatusUnprocessableEntity)
		assert.StringContains(t, string(body), `"first_name": "must be provided"`)
	})
}

func TestGetPatientHandler(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		app, _ := newTestApplication(t)
		srv := newTestServer(t, app.routes())
		defer srv.Close()

		statusCode, _, body := srv.getJSON(t, "/v1/patients/1")
		assert.Equal(t, statusCode, http.StatusOK)

		gotPatient := unmarshalPatient(t, body)
		assert.Equal(t, gotPatient.FirstName, "Ahmad")
		assert.Equal(t, gotPatient.LastName, "Abuziad")
		assert.Equal(t, gotPatient.Gender, "M")
		assert.Equal(t, gotPatient.DateOfBirth, time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC))
		assert.Equal(t, gotPatient.ID, 1)
		assert.Equal(t, gotPatient.CreatedAt.IsZero(), true)
	})

	t.Run("invalid id", func(t *testing.T) {
		app, _ := newTestApplication(t)
		srv := newTestServer(t, app.routes())
		defer srv.Close()

		statusCode, _, _ := srv.getJSON(t, "/v1/patients/invalid_id")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})

	t.Run("not found id 5", func(t *testing.T) {
		app, _ := newTestApplication(t)
		srv := newTestServer(t, app.routes())
		defer srv.Close()

		statusCode, _, _ := srv.getJSON(t, "/v1/patients/5")
		assert.Equal(t, statusCode, http.StatusNotFound)
	})
}

func unmarshalPatient(t *testing.T, b []byte) data.Patient {
	t.Helper()

	var jsRes struct {
		Patient data.Patient `json:"patient"`
	}
	err := json.Unmarshal(b, &jsRes)
	check(t, err)

	return jsRes.Patient
}
