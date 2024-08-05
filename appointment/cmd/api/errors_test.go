package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotFoundResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/path", nil)
	w := httptest.NewRecorder()

	app.notFoundResponse(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d but got %d", http.StatusNotFound, resp.StatusCode)
	}

	expectedBody := "the requested resource could not be found"
	if !containsString(w.Body.String(), expectedBody) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", expectedBody, w.Body.String())
	}
}

func TestMethodNotAllowedResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/path", nil)
	w := httptest.NewRecorder()

	app.methodNotAllowedResponse(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d but got %d", http.StatusMethodNotAllowed, resp.StatusCode)
	}

	expectedBody := "the GET method is not supported for this resourse"
	if !containsString(w.Body.String(), expectedBody) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", expectedBody, w.Body.String())
	}
}

func TestBadRequestResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/path", nil)
	w := httptest.NewRecorder()
	err := errors.New("bad request error")

	app.badRequestResponse(w, req, err)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, resp.StatusCode)
	}

	if !containsString(w.Body.String(), err.Error()) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", err.Error(), w.Body.String())
	}
}

func TestRateLimitExceededResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/path", nil)
	w := httptest.NewRecorder()

	app.rateLimitExceededResponse(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status code %d but got %d", http.StatusTooManyRequests, resp.StatusCode)
	}

	expectedBody := "rate limit exceeded"
	if !containsString(w.Body.String(), expectedBody) {
		t.Errorf("Expected response body to contain '%s' but got '%s'", expectedBody, w.Body.String())
	}
}

func containsString(s, substr string) bool {
	return strings.Contains(s, substr)
}
