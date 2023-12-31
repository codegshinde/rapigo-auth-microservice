package handlers

import (
	"net/http"
	"rapigo/internal/middleware"
	"rapigo/pkg/utils"
)

const (
	StatusUnauthorized = http.StatusUnauthorized
	BadJwtToken        = "Please sign in to access this information."
)

func Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)

	// Retrieve authenticated admin
	admin, err := middleware.GetAuthenticatedAdmin(r)
	if err != nil {
		utils.WriteErrorResponse(w, StatusUnauthorized, BadJwtToken)
		return
	}

	// Create JSON response
	responseJSON, err := utils.CreateJSONResponse(admin)
	if err != nil {
		utils.WriteErrorResponse(w, BadRequest, ErrorInternalServer)
		return
	}

	// Write the JSON response to the client
	w.Write(responseJSON)
}
