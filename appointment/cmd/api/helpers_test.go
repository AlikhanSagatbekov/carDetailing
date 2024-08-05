package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestReadJSON(t *testing.T) {
	app := &application{}

	type TestData struct {
		Message string `json:"message"`
	}

	body := `{"message": "Salem dos!"}`
	req := httptest.NewRequest("POST", "/data", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	var data TestData
	err := app.readJSON(httptest.NewRecorder(), req, &data)
	if err != nil {
		t.Errorf("readJSON failed: %s", err)
	}

	if data.Message != "Salem dos!" {
		t.Errorf("readJSON: expected message to be 'Salem dos!' but got %s", data.Message)
	}
}
