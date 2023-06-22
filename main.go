package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func LoadEnvFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	err := LoadEnvFromFile(".env")
	apiKey := os.Getenv("PALM_API_KEY")
	modelName := "text-bison-001"
	endpoint := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta2/models/%s:generateText?key=%s", modelName, apiKey)

	payload := `{
        "prompt": {"text": "what is palm api"},
        "temperature": 1.0,
        "candidate_count": 2
    }`

	jsonPayload := []byte(payload)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	fmt.Println(resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	fmt.Println(result["responses"])
}
