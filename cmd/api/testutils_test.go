package main

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) (*application, *bytes.Buffer) {
	t.Helper()

	var logBuf bytes.Buffer

	app := &application{
		logger: slog.New(slog.NewTextHandler(&logBuf, nil)),
	}

	return app, &logBuf
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	t.Helper()

	var srv *httptest.Server = httptest.NewServer(h)
	return &testServer{srv}
}

func (ts *testServer) postJSON(t *testing.T, urlPath string, body string) (int, http.Header, []byte) {
	t.Helper()

	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", bytes.NewReader([]byte(body)))
	check(t, err)

	rsBody := read(t, rs.Body)

	return rs.StatusCode, rs.Header, rsBody
}

func check(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

func read(t *testing.T, r io.ReadCloser) []byte {
	defer r.Close()
	data, err := io.ReadAll(r)
	check(t, err)
	return data
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("custom error")
}
