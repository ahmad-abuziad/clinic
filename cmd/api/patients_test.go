package main

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/assert"
	"github.com/ahmad-abuziad/clinic/internal/data"
)

func TestCreatePatientHandler(t *testing.T) {
	ts := newTestServer(t, routes())
	defer ts.Close()

	urlPath := "/v1/patient"

	t.Run("Valid request", func(t *testing.T) {
		patient := data.Patient{
			FirstName:   "Ahmad",
			LastName:    "Abuziad",
			DateOfBirth: time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC),
			Gender:      "M",
			Notes:       "Gluten Allergic",
		}

		statusCode, _, body := ts.postJSON(t, urlPath, patient)

		assert.Equal(t, statusCode, http.StatusCreated)

		gotPatient := unmarshalPatient(t, body)
		assert.Equal(t, gotPatient.FirstName, patient.FirstName)
		assert.Equal(t, gotPatient.LastName, patient.LastName)
		assert.Equal(t, gotPatient.DateOfBirth, patient.DateOfBirth)
		assert.Equal(t, gotPatient.Gender, gotPatient.Gender)
		assert.Equal(t, gotPatient.Notes, patient.Notes)
	})

	t.Run("Invalid JSON body", func(t *testing.T) {
		statusCode, _, body := ts.postJSON(t, urlPath, "{'invalid': json")

		assert.Equal(t, statusCode, http.StatusBadRequest)
		assert.StringContains(t, string(body), "body contains incorrect JSON type")
	})

	t.Run("Invalid passing extra fields", func(t *testing.T) {
		input := struct {
			data.Patient
			ExtraField string `json:"extra_field"`
		}{
			ExtraField: "extra data",
		}

		statusCode, _, body := ts.postJSON(t, urlPath, input)

		assert.Equal(t, statusCode, http.StatusBadRequest)
		assert.StringContains(t, string(body), `body contains unknown key "extra_field"`)
	})
}

func unmarshalPatient(t *testing.T, b []byte) data.Patient {
	t.Helper()

	var jsRes struct {
		Patient data.Patient `json:"patient"`
	}
	err := json.Unmarshal(b, &jsRes)
	if err != nil {
		t.Fatal(err)
	}

	return jsRes.Patient
}
