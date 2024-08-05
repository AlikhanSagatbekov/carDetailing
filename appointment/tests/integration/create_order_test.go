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
}
