package palm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Model struct {
	Name             string   `json:"name"`
	BaseModelId      string   `json:"baseModelId"`
	Version          string   `json:"version"`
	DisplayName      string   `json:"displayName"`
	Description      string   `json:"description"`
	InputTokenLimit  int      `json:"inputTokenLimit"`
	OutputTokenLimit int      `json:"outputTokenLimit"`
	SupportedMethods []string `json:"supportedGenerationMethods"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"topP"`
	TopK             int      `json:"topK"`
}

func ListModels() ([]Model, error) {
	err := loadEnvFromFile(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	apiKey := os.Getenv("PALM_API_KEY")
	endpoint := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta2/models/?key=%s", apiKey)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result map[string][]Model
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	if result["models"] == nil || len(result["models"]) == 0 {
		return nil, fmt.Errorf("no models found")
	}
	return result["models"], nil
}

func GetModel(name string) (Model, error) {
	err := loadEnvFromFile(".env")
	apiKey := os.Getenv("PALM_API_KEY")
	if err != nil {
		return Model{}, fmt.Errorf("error loading .env file: %w", err)
	}
	endpoint := fmt.Sprintf("%s/models/%s?key=%s", API_BASE_URL, name, apiKey)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return Model{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	var result Model
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return Model{}, fmt.Errorf("error decoding response: %w", err)
	}
	return result, nil
}
