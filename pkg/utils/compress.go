package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"rapigo/internal/models"
)

// Compress compresses an AdminResponse object to JSON format and applies gzip compression.
func Compress(adminResponse models.AdminResponse) ([]byte, error) {
	// Marshal the response object to JSON format.
	adminJSON, err := json.Marshal(adminResponse)
	if err != nil {
		return nil, err
	}

	// Create a buffer to write the compressed data to
	var compressedData bytes.Buffer

	// Create a gzip writer that writes to the buffer
	gzipWriter := gzip.NewWriter(&compressedData)

	// Write the JSON data to the gzip writer
	_, err = gzipWriter.Write(adminJSON)
	if err != nil {
		return nil, err
	}

	// Close the gzip writer to flush any remaining data
	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}
