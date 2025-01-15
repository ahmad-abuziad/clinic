package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestReadJSON(t *testing.T) {

	largeBody, err := os.ReadFile(path.Join("testdata", "largefile.json"))
	check(t, err)

	type dst struct {
		FieldName string `json:"field_name"`
	}

	t.Run("Valid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", strings.NewReader(`{"field_name": "field value"}`))

		d := dst{}
		err := readJSON(rr, r, &d)

		assert.NilError(t, err)
		assert.Equal(t, d.FieldName, "field value")
	})

	tests := []struct {
		name      string
		body      string
		wantError string
	}{
		{
			name:      "Invalid JSON body (at character)",
			body:      `{'key_surrounded_by_single_quotes':`,
			wantError: "body contains badly-formed JSON (at character 2)",
		},
		{
			name:      "Invalid JSON body",
			body:      `{"key_without_value":`,
			wantError: "body contains badly-formed JSON",
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
			body:      `{"field_name":1}`,
			wantError: `body contains incorrect JSON type for field "field_name"`,
		},
		{
			name:      "Incorrect field type",
			body:      `["field_name", "field_value"]`,
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
			rr := httptest.NewRecorder()
			body := strings.NewReader(tt.body)
			r := httptest.NewRequest("POST", "/", body)
			err := readJSON(rr, r, &dst{})
			assert.Equal(t, err.Error(), tt.wantError)
		})
	}

	t.Run("Panic when passing dst as a value", func(t *testing.T) {
		defer func() {
			recover()
		}()

		rr := httptest.NewRecorder()
		body := strings.NewReader("{}")
		r := httptest.NewRequest("POST", "/", body)
		readJSON(rr, r, dst{})
		t.Error("Should have panicked")
	})

	t.Run("Default error", func(t *testing.T) {
		errorReader := &errorReader{}
		req := httptest.NewRequest("POST", "/", errorReader)

		rr := httptest.NewRecorder()

		var dst any
		err := readJSON(rr, req, &dst)

		assert.NonNilError(t, err)
		assert.Equal(t, err.Error(), "custom error")
	})
}

func TestWriteJSON(t *testing.T) {

	t.Run("Valid", func(t *testing.T) {
		rr := httptest.NewRecorder()
		headers := make(http.Header)
		headers.Set("Location", "/location")

		err := writeJSON(rr, http.StatusCreated, envelope{"key": "value"}, headers)
		assert.NilError(t, err)

		rs := rr.Result()
		body := read(t, rs.Body)

		var m map[string]string
		err = json.Unmarshal(body, &m)
		check(t, err)

		assert.Equal(t, rs.Header.Get("Location"), "/location")
		assert.Equal(t, rs.StatusCode, http.StatusCreated)
		assert.Equal(t, m["key"], "value")
	})

	t.Run("Invalid - channels cannot be marshaled into JSON", func(t *testing.T) {
		rr := httptest.NewRecorder()
		got := writeJSON(rr, http.StatusCreated, envelope{"key": make(chan int)}, nil)

		assert.NonNilError(t, got)
	})

}
