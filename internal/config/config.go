package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadENVFiles reads environment variables from a file and returns them as a map.
func LoadENVFiles(filename string) (map[string]string, error) {
	env := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file '%s': %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			env[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file '%s': %w", filename, err)
	}

	return env, nil
}

func GetEnvVariable(variablesName string) string {
	env, err := LoadENVFiles(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return ""
	}
	return env[variablesName]
}
