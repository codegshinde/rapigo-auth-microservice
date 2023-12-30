// handlers/admin.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rapigo/service"
	"rapigo/utils"
)

// Credentials represents the JSON structure expected in the request body for login.
type Credentials struct {
	AdminId  string `json:"adminId"`
	Password string `json:"password"`
}

// Login handles the HTTP request for user login.
func Login(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response will be in JSON format.
	w.Header().Set("Content-Type", "application/json")

	// Decode the request body to extract credentials.
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	// Check for errors in decoding the request body.
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Retrieve admin information from the database based on the provided admin ID.
	admin, err := service.GetAdminByID(credentials.AdminId)
	if err != nil {
		// If there is an error fetching admin information, respond with a 400 Bad Request.
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Compare the provided password with the hashed password from the database.
	err = utils.ComparePassword([]byte(admin.Password), credentials.Password)
	if err != nil {
		// If passwords do not match, respond with a 400 Bad Request and an error message.
		fmt.Println("Passwords do not match:", err)
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Generate a JWT token for the authenticated admin.
	tokenString, err := utils.GenerateToken(credentials.AdminId)
	if err != nil {
		// If there is an error generating the token, respond with a 500 Internal Server Error.
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Prepare the response object containing the generated token and admin ID.
	responseObject := &utils.Response{
		Token:   tokenString,
		AdminId: credentials.AdminId,
	}

	// Convert the response object to JSON format.
	response, err := utils.CreateJSONResponse(responseObject)
	if err != nil {
		// If there is an error creating the JSON response, respond with a 500 Internal Server Error.
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write the JSON response to the HTTP response writer.
	w.Write([]byte(response))
}
