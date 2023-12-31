package handlers

import (
	"encoding/json"
	"net/http"
	"rapigo/internal/service"
	"rapigo/pkg/utils"
)

const (
	ContentTypeJSON      = "application/json"
	ErrorBadRequest      = "Invalid credentials"
	ErrorInternalServer  = "Internal Server Error"
	ErrorNoDocumentFound = "Admin not found"
)

// Credentials represents the structure for decoding JSON credentials in the login request.
type Credentials struct {
	AdminId  string `json:"adminId" bson:"adminId"`
	Password string `json:"password" bson:"password"`
}

// Login handles the login endpoint.
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeJSON)
	defer r.Body.Close()

	// Decode JSON credentials from the request body.
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Retrieve admin information based on the provided admin ID.
	admin, err := service.GetAdminByID(credentials.AdminId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, ErrorNoDocumentFound)
		return
	}

	// Compare the provided password with the hashed password stored in the database.
	err = utils.ComparePassword(admin.Password, credentials.Password)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, ErrorBadRequest)
		return
	}
	// Generate a JWT token for successful authentication.
	token, err := utils.GenerateToken(credentials.AdminId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, ErrorInternalServer)
		return
	}

	//admin without password
	adminWithoutPassword := admin.ConvertToAdminResponse(token)

	// Marshal the response object to JSON format.
	adminResponse, err := json.Marshal(adminWithoutPassword)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write the JSON response to the client.
	w.Write(adminResponse)
}
