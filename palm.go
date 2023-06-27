package palm

import (
	"bufio"
	"bytes"
	"net/http"
	"os"
	"strings"
)

func loadAPIKey(filepath string) (string, error) {
	apiKey := os.Getenv("PALM_API_KEY")
	if apiKey == "" {
		file, err := os.Open(filepath)
		if err != nil {
			return "", err
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
				apiKey = os.Getenv("PALM_API_KEY")
			}
		}

		if err := scanner.Err(); err != nil {
			return "", err
		}
	}
	return apiKey, nil

}

func makeRequest(endpoint, method string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
