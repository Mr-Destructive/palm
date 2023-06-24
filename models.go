package palm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// ListModels returns a list of all models from the Palm API
func ListModels() ([]Model, error) {
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	endpoint := fmt.Sprintf("%s/models/?key=%s", API_BASE_URL, apiKey)

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

// GetModel returns a single model by name from the Palm API
func GetModel(name string) (Model, error) {
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return Model{}, fmt.Errorf("error loading .env file: %w", err)
	}
	if !strings.HasPrefix(name, "models/") {
		name = "models/" + name
	}
	endpoint := fmt.Sprintf("%s/%s?key=%s", API_BASE_URL, name, apiKey)
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
