package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoverPanicMiddleware(t *testing.T) {
	app := &application{}

	handler := app.recoverPanic(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Code)
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	app := &application{}

	handler := app.rateLimit(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))

	// Simulate exceeding the rate limit
	for i := 0; i < 5; i++ { // Attempting 5 times to exceed the limit
		req, _ := http.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()
		handler.ServeHTTP(resp, req)
	}

	// Check if the rate limit was exceeded after 5 attempts
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	if resp.Code == http.StatusOK {
		t.Errorf("Expected status Too Many Requests, got %v", resp.Code)
	}
}
