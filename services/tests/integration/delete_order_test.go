package testing

import (
	"net/http"
	"testing"
)

func TestDeleteOrderHandler(t *testing.T) {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:4001/v1/services/1", nil)
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
