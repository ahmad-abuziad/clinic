package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"testing"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/assert"
	"github.com/ahmad-abuziad/clinic/internal/data"
)

func TestCreatePatientHandler(t *testing.T) {
	app := application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	urlPath := "/v1/patient"

	t.Run("201 Created", func(t *testing.T) {
		var (
			validFirstName   = "Ahmad"
			validLastName    = "Abuziad"
			validDateOfBirth = time.Date(1993, 6, 8, 0, 0, 0, 0, time.UTC)
			validGender      = "M"
			validNotes       = "Gluten Allergic"
		)

		reqBody := fmt.Sprintf(`
		{
			"first_name": %q,
			"last_name": %q,
			"date_of_birth": %q,
			"Gender": %q,
			"Notes": %q
		}`, validFirstName, validLastName, validDateOfBirth.Format(time.RFC3339), validGender, validNotes)

		statusCode, _, body := ts.postJSON(t, urlPath, reqBody)

		assert.Equal(t, statusCode, http.StatusCreated)

		gotPatient := unmarshalPatient(t, body)
		assert.Equal(t, gotPatient.FirstName, validFirstName)
		assert.Equal(t, gotPatient.LastName, validLastName)
		assert.Equal(t, gotPatient.DateOfBirth, validDateOfBirth)
		assert.Equal(t, gotPatient.Gender, validGender)
		assert.Equal(t, gotPatient.Notes, validNotes)
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
