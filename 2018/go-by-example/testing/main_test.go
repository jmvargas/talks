package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct{ in, out string }{
		{"jesusmanuel.vargas@privalia.com", "Hello jesusmanuel.vargas! 👋"},
		{"foo.bar@other.com", "Sorry, I don't know you"},
	}

	for _, c := range cases {
		req, _ := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)
		rec := httptest.NewRecorder()
		Handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Unexpected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("Unexpected body: %s", rec.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		req, _ := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/jesusmanuel.vargas@privalia.com",
			nil,
		)
		rec := httptest.NewRecorder()
		Handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Unexpected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "Hello jesusmanuel.vargas! 👋") {
			b.Errorf("Unexpected body: %s", rec.Body.String())
		}
	}
}
