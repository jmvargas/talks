package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	body := strings.NewReader(`{"name":"Jesús","age":26}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:8080/",
		body,
	)

	rec := httptest.NewRecorder()
	Handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Unexpected status 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Jesús, you can drink beer 🍺") {
		t.Errorf("Unexpected body: %s", rec.Body.String())
	}
}

func TestHandlerFail(t *testing.T) {
	body := strings.NewReader(`{"name":"Jesús","age":-1}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"http://localhost:8080/",
		body,
	)
	rec := httptest.NewRecorder()
	Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Unexpected status 400, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Error: age lower than zero 💥") {
		t.Errorf("Unexpected body: %s", rec.Body.String())
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		body := strings.NewReader(`{"name":"Jesús","age":26}`)

		req := httptest.NewRequest(
			http.MethodPost,
			"http://localhost:8080/",
			body,
		)
		rec := httptest.NewRecorder()
		Handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Unexpected status 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "Jesús, you can drink beer 🍺") {
			b.Errorf("Unexpected body: %s", rec.Body.String())
		}
	}
}
