package palm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)


type ResponseText struct {
	Candidates     []TextCompletion `json:"candidates"`
	Filters        []ContentFilter  `json:"filters"`
	SafetyFeedback []SafetyFeedback `json:"safetyFeedback"`
}

func GenerateText(model string, params PromptConfig) (string, error) {

	if model == "" {
		model = "text-bison-001"
	}
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return "", err
	}
	endpoint := fmt.Sprintf("%s/models/%s:generateText?key=%s", API_BASE_URL, model, apiKey)

	if params.CandidateCount <= 0 {
		params.CandidateCount = 1
	}
	if params.MaxOutputTokens <= 0 {
		params.MaxOutputTokens = 100
	}
	if params.TopK <= 0 {
		params.TopK = 1
	}

	jsonMessagePrompt, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	payload := string(jsonMessagePrompt)
	jsonPayload := []byte(payload)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var result ResponseText
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if len(result.Candidates) > 0 {
		return result.Candidates[0].Output, nil
	}
	return "", fmt.Errorf("error fetching response data")
}

type ResponseMessage struct {
	Candidates []Message       `json:"candidates"`
	Messages   []Message       `json:"messages"`
	Filters    []ContentFilter `json:"filters"`
}

func GenerateMessage(messages MessagePrompt, params map[string]string) (string, error) {
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return "", err
	}
	endpoint := fmt.Sprintf("%s/models/%s:generateMessage?key=%s", API_BASE_URL, params["model"], apiKey)
	jsonMessagePrompt, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}
	payload := `{
	    "prompt": ` + string(jsonMessagePrompt) + `,
        "temperature": 0.2,
        "candidate_count": 1
    }`
	jsonPayload := []byte(payload)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var result ResponseMessage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if len(result.Candidates) > 0 {
		return result.Candidates[0].Content, nil
	}
	return "", fmt.Errorf("error fetching response candidates")
}

type ResponseEmbed struct {
	Embedding Embedding `json:"embedding"`
}

func EmbedText(text string) (ResponseEmbed, error) {
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return ResponseEmbed{}, err
	}
	endpoint := fmt.Sprintf("%s/%s:embedText?key=%s", API_BASE_URL, "models/embedding-gecko-001", apiKey)
	jsonMessagePrompt, err := json.Marshal(text)
	if err != nil {
		return ResponseEmbed{}, err
	}
	payload := `{
        "text": ` + string(jsonMessagePrompt) + `
    }`
	jsonPayload := []byte(payload)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return ResponseEmbed{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseEmbed{}, err
	}

	var result ResponseEmbed
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ResponseEmbed{}, err
	}
	return result, nil
}
