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

	appointment := data.Appointment{
		CustomerID: 1,
		ServiceID:  1,
		Date:       "08.06.2024",
		Status:     "Started",
	}

	jsonAppointment, _ := json.Marshal(appointment)

	req, err := http.NewRequest("PUT", "http://localhost:4000/v1/appointments/3", bytes.NewBuffer(jsonAppointment))
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
