package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AlikhanSagatbekov/adv11/internal/data"
)

func TestCreateOrderHandler(t *testing.T) {
	client := &http.Client{}

	service := data.Service{
		Name:        "Cleaning",
		Description: "Cleaning outdoor and indoor",
		Price:       10000,
	}

	jsonServices, _ := json.Marshal(service)

	req, err := http.NewRequest("POST", "http://localhost:4001/v1/services", bytes.NewBuffer(jsonServices))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}
