package main

import (
	"bytes"
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
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	rsBody, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, rsBody
}
