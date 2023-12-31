// login_test.go
package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rapigo/internal/handlers"
	"rapigo/internal/models"
	"testing"
)

func BenchmarkLogin(b *testing.B) {
	// Prepare a sample request with a known payload
	requestPayload := []byte(`{"adminId": "sampleID", "password": "samplePassword"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(requestPayload))
	if err != nil {
		b.Fatalf("Failed to create request: %v", err)
	}

	for i := 0; i < b.N; i++ {
		// Create a response recorder to capture the response
		w := httptest.NewRecorder()

		// Run the actual handler
		handlers.Login(w, req)

		// Check the status code and response body if needed
		if w.Code != http.StatusOK {
			b.Fatalf("Unexpected status code: got %v, want %v", w.Code, http.StatusOK)
		}

		// Unmarshal the JSON response if needed
		var adminResponse models.AdminResponse
		if err := json.Unmarshal(w.Body.Bytes(), &adminResponse); err != nil {
			b.Fatalf("Failed to unmarshal JSON response: %v", err)
		}
	}
}
