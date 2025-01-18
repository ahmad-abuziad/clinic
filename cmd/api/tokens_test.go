package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestCreateAuthenticationTokenHandler(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		app, _ := newTestApplication(t)
		body := `
		{
			"email": "ahmad@example.com",
			"password": "pa55word"
		}
		`
		rr := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/v1/tokens/authentication", strings.NewReader(body))

		app.createAuthenticationTokenHandler(rr, r)

		rs := rr.Result()
		rsBody := read(t, rs.Body)

		var response struct {
			Token string `json:"token"`
		}

		json.Unmarshal(rsBody, &response)
		//assert.Equal(t, response.Token, "token")
		//assert.Equal(t, len(response.Token), 26)
	})

	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{
			name:     "invalid JSON",
			body:     `{"email":}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "invalid email",
			body:     `{"email":"invalid_email.com"}`,
			wantCode: http.StatusUnprocessableEntity,
		},
		{
			name:     "invalid password",
			body:     `{"password":"short"}`,
			wantCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app, _ := newTestApplication(t)
			rr := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/v1/tokens/authentication", strings.NewReader(tt.body))

			app.createAuthenticationTokenHandler(rr, r)

			rs := rr.Result()

			assert.Equal(t, rs.StatusCode, tt.wantCode)
		})
	}
}
