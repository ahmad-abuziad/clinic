package main

import (
	"bytes"
	"database/sql"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/ahmad-abuziad/clinic/internal/data"
)

func newTestApplication(t *testing.T) (*application, *bytes.Buffer) {
	t.Helper()

	var logBuf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logBuf, nil))
	db := newTestDB(t)
	app := &application{
		logger: logger,
		models: data.NewModels(db),
		errors: newHTTPErrors(logger),
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

func (ts *testServer) getJSON(t *testing.T, urlPath string) (int, http.Header, []byte) {
	t.Helper()

	rs, err := ts.Client().Get(ts.URL + urlPath)
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

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("TEST_CLINIC_DB_DSN"))
	if err != nil {
		t.Fatal(err)
	}

	migPath := path.Join("..", "..", "migrations")
	c, err := os.ReadDir(migPath)
	if err != nil {
		db.Close()
		t.Fatal(err)
	}

	ups := []string{}
	downs := []string{}
	for _, entry := range c {
		if strings.HasSuffix(entry.Name(), ".up.sql") {
			ups = append(ups, path.Join(migPath, entry.Name()))
		}
		if strings.HasSuffix(entry.Name(), ".down.sql") {
			downs = append(downs, path.Join(migPath, entry.Name()))
		}
	}
	ups = append(ups, path.Join("testdata", "setup.sql"))
	downs = append(downs, path.Join("testdata", "teardown.sql"))

	for _, path := range ups {
		script, err := os.ReadFile(path)
		if err != nil {
			db.Close()
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			db.Close()
			t.Fatal(err)
		}
	}

	t.Cleanup(func() {
		defer db.Close()

		for i := len(downs) - 1; i >= 0; i-- {
			path := downs[i]
			script, err := os.ReadFile(path)
			if err != nil {
				db.Close()
				t.Fatal(err)
			}
			_, err = db.Exec(string(script))
			if err != nil {
				db.Close()
				t.Fatal(err)
			}
		}
	})

	return db
}
