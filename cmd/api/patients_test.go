package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ahmad-abuziad/clinic/internal/assert"
	"github.com/ahmad-abuziad/clinic/internal/data"
)

func TestCreatePatientHandler(t *testing.T) {
	ts := newTestServer(t, routes())
	defer ts.Close()

	urlPath := "/v1/patient"

	largeBody, err := os.ReadFile("largefile.json")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid request", func(t *testing.T) {
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

	tests := []struct {
		name      string
		body      string
		wantError string
	}{
		{
			name:      "Invalid JSON body",
			body:      "{'invalid': json",
			wantError: "body contains badly-malformed JSON (at character 2)",
		},
		{
			name: "Invalid passing extra fields",
			body: `
				{
					"extra_field": "extra data"
				}
			`,
			wantError: `body contains unknown key "extra_field"`,
		},
		{
			name:      "Empty body",
			body:      "",
			wantError: "body must not be empty",
		},
		{
			name:      "Incorrect field type",
			body:      `{"first_name":1}`,
			wantError: `body contains incorrect JSON type for field "first_name"`,
		},
		{
			name:      "Incorrect field type",
			body:      `["first_name", "ahmad"]`,
			wantError: `body contains incorrect JSON type (at character 1)`,
		},
		{
			name:      "Two JSON objects in body",
			body:      "{}{}",
			wantError: "body must only contain a single JSON value",
		},
		{
			name:      "body > 1MB",
			body:      string(largeBody),
			wantError: "body must not be larger than 1048576 bytes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, body := ts.postJSON(t, urlPath, tt.body)

			assert.Equal(t, statusCode, http.StatusBadRequest)
			assert.StringContains(t, string(body), tt.wantError)
		})
	}
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
