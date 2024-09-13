package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func noopHandler(c *Ctx) error {
	return nil
}

func BenchmarkSomeRequest(b *testing.B) {
	l := New(true)
	l.Get("/", noopHandler)

	for i := 0; i < b.N; i++ {
		doRequest(nil, "GET", "/", nil, l)
	}
}

func doRequest(t *testing.T, method, route string, body io.Reader, l *Launch) {
	r, err := http.NewRequest(method, route, body)
	if err != nil {
		t.Fatal(err)
	}
	rw := httptest.NewRecorder()
	l.router.ServeHTTP(rw, r)
}
