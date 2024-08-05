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

	appointment := data.Appointment{
		CustomerID: 1,
		ServiceID:  12,
		Date:       "08.06.2024",
		Status:     "Not started",
	}

	jsonAppointment, _ := json.Marshal(appointment)

	req, err := http.NewRequest("POST", "http://localhost:4000/v1/appointments", bytes.NewBuffer(jsonAppointment))
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

	updatedAppointment := data.Appointment{
		CustomerID: 1,
		ServiceID:  12,
		Date:       "08.06.2024",
		Status:     "Started",
	}

	jsonUpdatedAppointment, _ := json.Marshal(updatedAppointment)

	req1, err := http.NewRequest("PUT", "http://localhost:4000/v1/appointments/2", bytes.NewBuffer(jsonUpdatedAppointment))
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

	req2, err := http.NewRequest("DELETE", "http://localhost:4000/v1/appointments/2", nil)
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
