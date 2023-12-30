package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func SpeedHandler(w http.ResponseWriter, r *http.Request) {
	// Mark the start time
	startTime := time.Now()

	// Simulate some work (optional)
	time.Sleep(50 * time.Millisecond) // Example: simulate 50ms of processing

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	// Prepare the response
	response := fmt.Sprintf("Speed test response took %s", elapsedTime)

	// Send the response
	w.Write([]byte(response))
}


