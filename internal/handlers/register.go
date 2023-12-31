package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rapigo/internal/models"
	"rapigo/internal/service"
	"rapigo/pkg/utils"

	"time"
)

const (
	MobileField = "mobile"
	EmailField  = "email"

	FirstName = "FirstName"
	LastName  = "LastName"
	Email     = "Email"
	Mobile    = "Mobile"
	Password  = "Password"

	BadRequest        = http.StatusBadRequest
	InvalidJSONFormat = "Invalid JSON format"
	RequiredField     = "%s is required"
	MobileExists      = "mobile number already exists"
	EmailExists       = "email address already exists"
	UnknownField      = "unknown field name"
)

// validateAdmin checks if required fields in Admin are empty.
func ValidateAdmin(admin models.AdminInput, w http.ResponseWriter) error {
	requiredFields := []string{FirstName, LastName, Email, Mobile, Password}
	for _, field := range requiredFields {
		if fieldValue := utils.GetField(admin, field); fieldValue == "" {
			return errors.New(field + RequiredField)
		}
	}
	return nil
}

// checkExistingAdmin checks if a admin with the provided field value already exists.
func CheckExistingAdmin(field, value string) error {
	response, _ := service.GetFindMyField(field, value)

	switch field {
	case MobileField:
		if response.Mobile != "" {
			return errors.New(MobileExists)
		}
	case EmailField:
		if response.Email != "" {
			return errors.New(EmailExists)
		}
	default:
		return errors.New(UnknownField)
	}

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate that the response will be in JSON format.
	w.Header().Set("Content-Type", ContentTypeJSON)

	var admin models.AdminInput
	err := json.NewDecoder(r.Body).Decode(&admin)

	if err != nil {
		utils.WriteErrorResponse(w, BadRequest, InvalidJSONFormat)
		return
	}

	// Validate required fields in Admin.
	if err := ValidateAdmin(admin, w); err != nil {
		utils.WriteErrorResponse(w, BadRequest, err.Error())
		return
	}

	// Check if the mobile number already exists.
	if err := CheckExistingAdmin(MobileField, admin.Mobile); err != nil {
		utils.WriteErrorResponse(w, BadRequest, MobileExists)
		return
	}

	// Check if the email already exists.
	if err := CheckExistingAdmin(EmailField, admin.Email); err != nil {
		utils.WriteErrorResponse(w, BadRequest, EmailExists)
		return
	}

	// Hash the password before storing it.
	password, err := utils.HashPassword(admin.Password)
	if err != nil {
		// If there is an error hashing the password, respond with a 400 Bad Request.
		utils.WriteErrorResponse(w, BadRequest, err.Error())
		return
	}

	adminId := utils.GenerateUniqId(admin.FirstName, admin.LastName)

	// Update the admin struct with the hashed password and current timestamp.
	admin.Password = password
	admin.AdminId = adminId
	admin.CreatedAt = time.Now()

	// Insert the admin information into the database.
	result, err := service.InsertOne(admin)
	if err != nil {
		// If there is an error inserting admin information, respond with a 400 Bad Request.
		utils.WriteErrorResponse(w, BadRequest, err.Error())
		return
	}

	// Convert the result to a JSON response.
	response, _ := utils.CreateJSONResponse(result)
	w.Write([]byte(response))
}
