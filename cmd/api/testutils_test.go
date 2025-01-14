package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

	defer rs.Body.Close()
	rsBody, err := io.ReadAll(rs.Body)
	check(t, err)

	return rs.StatusCode, rs.Header, rsBody
}

func check(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("custom error")
}
