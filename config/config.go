package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadENVFiles(filename string) (map[string]string, error) {
	env := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return env, nil
}

func GetEnvVariables(variablesName string) string {
	env, err := LoadENVFiles(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return ""
	}
	return env[variablesName]
}
