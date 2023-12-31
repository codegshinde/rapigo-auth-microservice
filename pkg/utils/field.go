package utils

import "rapigo/internal/models"

// GetField gets the value of a specified field in the Admin struct.
func GetField(admin models.AdminInput, fieldName string) string {
	switch fieldName {
	case "AdminId":
		return admin.AdminId
	case "FirstName":
		return admin.FirstName
	case "LastName":
		return admin.LastName
	case "AdminRole":
		return admin.AdminRole
	case "AdminPosition":
		return admin.AdminPosition
	case "Email":
		return admin.Email
	case "Mobile":
		return admin.Mobile
	case "Password":
		return admin.Password
	case "CreatedAt":
		return admin.CreatedAt.String()
	default:
		return ""
	}
}
