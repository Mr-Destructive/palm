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

func GenerateMessage(messageConfig MessageConfig) (*ResponseMessage, error) {
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return nil, err
	}
	model := "chat-bison-001"
	messages := messageConfig.Prompt
	endpoint := fmt.Sprintf("%s/models/%s:generateMessage?key=%s", API_BASE_URL, model, apiKey)
	jsonMessagePrompt, err := json.Marshal(messages)
	if err != nil {
		return nil, err
	}
	payload := `{
	    "prompt": ` + string(jsonMessagePrompt) + `,
        "temperature": 0.2,
        "candidate_count": 1
    }`
	jsonPayload := []byte(payload)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result ResponseMessage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	if len(result.Candidates) > 0 {
		return &result, nil
	}
	return nil, fmt.Errorf("error fetching response candidates")
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

type ChatResponse struct {
	Candidates     []Message       `json:"candidates"`
	Filters        []ContentFilter `json:"filters"`
	Messages       []MessagePrompt `json:"messages"`
	Model          string          `json:"model"`
	Context        string          `json:"context"`
	Examples       []Example       `json:"examples"`
	Temperature    float64         `json:"temperature"`
	CandidateCount int             `json:"candidate_count"`
	TopK           int             `json:"top_k"`
	TopP           float64         `json:"top_p"`
	Last           string          `json:"last"`
}

func (c *ChatResponse) GetLast() string {
	if len(c.Candidates) > 0 {
		return c.Candidates[0].Content
	}
	return ""
}

func Chat(config ChatConfig) (ChatResponse, error) {
	var msg string
	if len(config.Messages) == 0 && config.Prompt.Text != "" {
		msg = config.Prompt.Text
	} else {
		msg = config.Messages[0].Content
	}
	msgConfig := MessageConfig{
		Prompt: MessagePrompt{
			Messages: []Message{Message{Content: msg}},
		},
		Temperature:    config.Temperature,
		CandidateCount: config.CandidateCount,
		TopK:           config.TopK,
		TopP:           config.TopP,
	}
	msgResp, err := GenerateMessage(msgConfig)
	if err != nil {
		return ChatResponse{}, err
	}
	resp := *&msgResp
	// convert the msg to MessagePrompt
	msgPrompt := []MessagePrompt{}
	for i, m := range resp.Candidates {
		resp.Messages[i].Content = m.Content
		message := Message{Content: m.Content}
		msgPrompt = append(msgPrompt, MessagePrompt{Messages: []Message{message}})
	}

	chatResp := ChatResponse{
		Candidates:     resp.Candidates,
		Filters:        resp.Filters,
		Messages:       msgPrompt,
		Model:          config.Model,
		Context:        config.Context,
		Examples:       config.Examples,
		Temperature:    config.Temperature,
		CandidateCount: config.CandidateCount,
		TopK:           config.TopK,
		TopP:           config.TopP,
	}
	chatResp.Last = chatResp.GetLast()
	return chatResp, nil

}

func (c *ChatResponse) Reply(message Message) {
	msgConfig := MessageConfig{
		Prompt: MessagePrompt{
			Messages: []Message{message},
		},
		Temperature:    c.Temperature,
		CandidateCount: c.CandidateCount,
		TopK:           c.TopK,
		TopP:           c.TopP,
	}
	msgResp, err := GenerateMessage(msgConfig)
	if err != nil {
		return
	}
	resp := *msgResp
	c.Messages = append(c.Messages, MessagePrompt{Messages: []Message{message}})
	c.Candidates = append(c.Candidates, resp.Candidates...)
	c.Last = resp.Candidates[0].Content
}

func ChatPrompt(prompt string) (ChatResponse, error) {
	chatConfig := ChatConfig{Prompt: TextPrompt{Text: prompt}}
	resp, err := Chat(chatConfig)
	if err != nil {
		return ChatResponse{}, err
	}
	return resp, nil
}
