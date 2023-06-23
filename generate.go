package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Candidates []struct {
		Output string `json:"output"`
	} `json:"candidates"`
}

type TextPrompt struct {
	Text string `json:"text"`
}

type SafetySetting struct {
	StopSequence string `json:"stopSequence"`
}

type PromptConfig struct {
	Prompt          TextPrompt      `json:"prompt"`
	SafetySettings  []SafetySetting `json:"safetySettings"`
	StopSequences   []string        `json:"stopSequences"`
	Temperature     float64         `json:"temperature"`
	CandidateCount  int             `json:"candidateCount"`
	MaxOutputTokens int             `json:"maxOutputTokens"`
	TopP            float64         `json:"topP"`
	TopK            int             `json:"topK"`
}

type Message struct {
	Author           string           `json:"author"`
	Content          string           `json:"content"`
	CitationMetadata CitationMetadata `json:"citationMetadata,omitempty"`
}

type MessagePrompt struct {
	Context  string    `json:"context"`
	Examples []Example `json:"examples,omitempty"`
	Messages []Message `json:"messages"`
}

type Example struct {
	Input  Message `json:"input,omitempty"`
	Output Message `json:"output,omitempty"`
}

type CitationMetadata struct {
	CitationSource []CitationSource `json:"citationSource,omitempty"`
}

type CitationSource struct {
	StartIndex int    `json:"startIndex,omitempty"`
	EndIndex   int    `json:"endIndex,omitempty"`
	Uri        string `json:"uri,omitempty"`
	License    string `json:"license,omitempty"`
}

type BlockedReason int

const (
	BLOCKED_REASON_UNSPECIFIED BlockedReason = iota
	SAFETY
	OTHER
)

func (r BlockedReason) String() string {
	switch r {
	case BLOCKED_REASON_UNSPECIFIED:
		return "BLOCKED_REASON_UNSPECIFIED"
	case SAFETY:
		return "SAFETY"
	case OTHER:
		return "OTHER"
	}
	return ""
}

type ContentFilter struct {
	Reason  BlockedReason `json:"reason"`
	Message string        `json:"message"`
}

func GenerateText(model string, params PromptConfig) (string, error) {

	if model == "" {
		model = "text-bison-001"
	}
	err := LoadEnvFromFile(".env")
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
	err := LoadEnvFromFile(".env")
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("PALM_API_KEY")
	endpoint := fmt.Sprintf("%s/models/%s:generateMessage?key=%s", API_BASE_URL, params["model"], apiKey)
	jsonMessagePrompt, err := json.Marshal(messages)
	if err != nil {
		return "", err
	}
	jsonMessagePrompt = []byte(strings.ReplaceAll(string(jsonMessagePrompt), "null", "\"\""))
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
		fmt.Println(err)
	}

	var result ResponseMessage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Candidates[0].Content, nil
}
