package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/assert"
)

func TestErrorResponse(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		app, _ := newTestApplication(t)

		responseRecorder := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		app.errorResponse(responseRecorder, r, http.StatusBadRequest, "")

		rs := responseRecorder.Result()
		assert.Equal(t, rs.StatusCode, http.StatusBadRequest)
	})

	t.Run("Invalid", func(t *testing.T) {
		app, logBuf := newTestApplication(t)

		rr := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		app.errorResponse(rr, r, http.StatusBadRequest, make(chan int))

		rs := rr.Result()
		assert.Equal(t, rs.StatusCode, http.StatusInternalServerError)

		log := logBuf.String()
		assert.StringContains(t, log, `level=ERROR msg="json: unsupported type: chan int" method=GET uri=/`)
	})
}

func TestLogError(t *testing.T) {
	app, logBuf := newTestApplication(t)

	r := httptest.NewRequest("GET", "/", nil)

	app.logError(r, errors.New("error message"))
	log := logBuf.String()

	assert.StringContains(t, log, `level=ERROR msg="error message" method=GET uri=/`)
}

func TestResponses(t *testing.T) {
	app, _ := newTestApplication(t)

	t.Run("badRequest", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		e := errors.New("error message")

		app.badRequest(rr, r, e)

		rs := rr.Result()
		body := read(t, rs.Body)

		assert.Equal(t, rs.StatusCode, http.StatusBadRequest)
		assert.StringContains(t, string(body), `"error": "error message"`)
	})

	t.Run("failedValidationResponse", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)
		errors := map[string]string{
			"field": "this field got an error",
		}

		app.failedValidationResponse(rr, r, errors)

		rs := rr.Result()
		body := read(t, rs.Body)

		assert.Equal(t, rs.StatusCode, http.StatusUnprocessableEntity)
		assert.StringContains(t, string(body), `"field": "this field got an error"`)
	})

	t.Run("notFoundResponse", func(t *testing.T) {
		rr := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", nil)

		app.notFoundResponse(rr, r)

		rs := rr.Result()
		body := read(t, rs.Body)

		assert.Equal(t, rs.StatusCode, http.StatusNotFound)
		assert.StringContains(t, string(body), `"error": "the requested resource could not be found`)
	})
}
