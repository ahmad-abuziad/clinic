package main

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestReadJSON(t *testing.T) {

	largeBody, err := os.ReadFile("largefile.json")
	if err != nil {
		t.Fatal(err)
	}
	var dst struct {
		FieldName string `json:"field_name"`
	}

	tests := []struct {
		name      string
		body      string
		wantError string
	}{
		{
			name:      "Invalid JSON body",
			body:      "{'invalid': json",
			wantError: "body contains badly-formed JSON (at character 2)",
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
			err := readJSON(rr, r, &dst)
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
		readJSON(rr, r, dst)
		t.Error("Should have panicked")
	})
}
