package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileInfo struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	IsDirectory  bool   `json:"isDirectory"`
	Length       int64  `json:"length"`
	LastModified string `json:"lastModified"`
}

func WriteEnvToFile(key, value string) error {
	outputFile, err := os.OpenFile(os.Getenv("DRONE_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer outputFile.Close()

	_, err = fmt.Fprintf(outputFile, "%s=%s\n", key, value)
	if err != nil {
		return fmt.Errorf("failed to write to env: %w", err)
	}

	return nil
}

func main() {
	// Sample JSON data representing file details
	files := []FileInfo{
		{
			Name:         "text1.txt",
			Path:         "/path/to/text1.txt",
			IsDirectory:  false,
			Length:       12345,
			LastModified: "2024-09-05T12:34:56Z",
		},
		{
			Name:         "text2.txt",
			Path:         "/path/to/text2.txt",
			IsDirectory:  false,
			Length:       67890,
			LastModified: "2024-09-05T12:35:12Z",
		},
		{
			Name:         "text3.txt",
			Path:         "/path/to/text3.txt",
			IsDirectory:  false,
			Length:       54321,
			LastModified: "2024-09-05T12:36:15Z",
		},
	}

	// Convert the files array to a JSON string
	jsonData, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}

	// Write the JSON string to the environment variable using WriteEnvToFile
	err = WriteEnvToFile("FILES_INFO", string(jsonData))
	if err != nil {
		fmt.Printf("Error writing to env: %v\n", err)
		return
	}

	fmt.Println("Successfully wrote FILES_PATH to environment")
}
