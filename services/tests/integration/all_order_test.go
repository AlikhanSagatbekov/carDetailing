package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/AlikhanSagatbekov/adv11/internal/data"
)

func TestCreateEditDeleteOrderHandler(t *testing.T) {
	client := &http.Client{}

	service := data.Service{
		Name:        "Cleaning",
		Description: "Cleaning outdoor and indoor",
		Price:       10000,
	}

	jsonService, _ := json.Marshal(service)

	req, err := http.NewRequest("POST", "http://localhost:4001/v1/services", bytes.NewBuffer(jsonService))
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

	updatedService := data.Service{
		Name:        "Cleaning",
		Description: "Cleaning outdoor and indoor",
		Price:       10000,
	}
	jsonUpdatedService, _ := json.Marshal(updatedService)

	req1, err := http.NewRequest("PUT", "http://localhost:4001/v1/services/2", bytes.NewBuffer(jsonUpdatedService))
	if err != nil {
		t.Fatal(err)
	}

	resp1, err := client.Do(req1)
	if err != nil {
		t.Fatal(err)
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	req2, err := http.NewRequest("DELETE", "http://localhost:4001/v1/appointments/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp2, err := client.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
