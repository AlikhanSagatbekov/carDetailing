package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AlikhanSagatbekov/carDetailing/internal/data"
)

func TestEditOrderHandler(t *testing.T) {
	client := &http.Client{}

	service := data.Service{
		Name:        "Cleaning",
		Description: "Cleaning outdoor and indoor",
		Price:       12000,
	}

	jsonServices, _ := json.Marshal(service)

	req, err := http.NewRequest("PUT", "http://localhost:4001/v1/services/1", bytes.NewBuffer(jsonServices))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
