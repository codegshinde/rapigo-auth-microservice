package utils

import "encoding/json"

// CreateJSONResponse creates a JSON response from the provided data.
func CreateJSONResponse(data interface{}) ([]byte, error) {
	responseJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return responseJSON, nil
}
