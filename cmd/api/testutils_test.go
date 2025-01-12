package main

import (
	"bytes"
	"encoding/json"
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

func (ts *testServer) postJSON(t *testing.T, urlPath string, data any) (int, http.Header, []byte) {
	t.Helper()

	js, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewReader(js)

	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", body)
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
