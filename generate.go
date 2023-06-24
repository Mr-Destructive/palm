package palm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Candidates []struct {
		Output string `json:"output"`
	} `json:"candidates"`
}

type ResponseText struct {
	Candidates     []TextCompletion `json:"candidates"`
	Filters        []ContentFilter  `json:"filters"`
	SafetyFeedback []SafetyFeedback `json:"safetyFeedback"`
}

func GenerateText(model string, params PromptConfig) (string, error) {

	if model == "" {
		model = "text-bison-001"
	}
	err := loadEnvFromFile(".env")
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("PALM_API_KEY")
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
		fmt.Println(err)
	}

	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Candidates[0].Output, nil
}

type ResponseMessage struct {
	Candidates []Message       `json:"candidates"`
	Messages   []Message       `json:"messages"`
	Filters    []ContentFilter `json:"filters"`
}

func GenerateMessage(messages MessagePrompt, params map[string]string) (string, error) {
	err := loadEnvFromFile(".env")
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("PALM_API_KEY")
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
	return result.Candidates[0].Content, nil
}

type ResponseEmbed struct {
	Embedding []Embedding `json:"embedding"`
}

func EmbedText(text string) (ResponseEmbed, error) {
	err := loadEnvFromFile(".env")
	if err != nil {
		return ResponseEmbed{}, err
	}
	apiKey := os.Getenv("PALM_API_KEY")
	endpoint := fmt.Sprintf("%s/models/%s:embedText?key=%s", API_BASE_URL, "text-bison-001", apiKey)
	jsonMessagePrompt, err := json.Marshal(text)
	if err != nil {
		return ResponseEmbed{}, err
	}
	jsonPayload := []byte(jsonMessagePrompt)
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
