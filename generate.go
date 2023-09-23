package palm

import (
	"encoding/json"
	"fmt"
)

type ResponseText struct {
	Candidates []TextCompletion `json:"candidates"`
}

type ResponseChat struct {
	Candidates     []TextCompletion `json:"candidates"`
	Filters        []ContentFilter  `json:"filters"`
	SafetyFeedback []SafetyFeedback `json:"safetyFeedback"`
}

func GenerateText(model, apiKey string, params PromptConfig) (string, error) {
	if model == "" {
		model = "text-bison-001"
	}
	if apiKey == "" {
		key, err := loadAPIKey(".env")
		if err != nil {
			return "", err
		}
		apiKey = key
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

	resp, err := makeRequest(endpoint, "POST", jsonPayload)
	if err != nil {
		return "", err
	}

	var result ResponseChat
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

func GenerateMessage(apiKey string, messageConfig MessageConfig) (*ResponseMessage, error) {
	if apiKey == "" {
		key, err := loadAPIKey(".env")
		if err != nil {
			return nil, err
		}
		apiKey = key
	}
	messages := messageConfig.Prompt
	endpoint := fmt.Sprintf("%s/%s:generateMessage?key=%s", API_BASE_URL, CHAT_MODEL, apiKey)
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
	resp, err := makeRequest(endpoint, "POST", jsonPayload)
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
	endpoint := fmt.Sprintf("%s/%s:embedText?key=%s", API_BASE_URL, EMBED_MODEL, apiKey)
	jsonMessagePrompt, err := json.Marshal(text)
	if err != nil {
		return ResponseEmbed{}, err
	}
	payload := `{
        "text": ` + string(jsonMessagePrompt) + `
    }`
	jsonPayload := []byte(payload)
	resp, err := makeRequest(endpoint, "POST", jsonPayload)
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
	Messages       []Message       `json:"messages"`
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
	message := Message{Content: msg, Author: "user"}
	if len(config.Messages) == 0 && config.Prompt.Text != "" {
		msg = config.Prompt.Text
		message.Content = msg
		config.Messages = append(config.Messages, message)
	} else {
		msg = config.Messages[len(config.Messages)-1].Content
	}
	msgConfig := MessageConfig{
		Prompt: MessagePrompt{
			//Messages: []Message{Message{Content: msg, Author: ""}},
			Messages: config.Messages,
		},
		Temperature:    config.Temperature,
		CandidateCount: config.CandidateCount,
		TopK:           config.TopK,
		TopP:           config.TopP,
	}
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return ChatResponse{}, err
	}
	msgResp, err := GenerateMessage(apiKey, msgConfig)
	if err != nil {
		return ChatResponse{}, err
	}
	resp := *&msgResp
	m := resp.Candidates[0]
	//resp.Messages[len(resp.Messages)-1].Content = m.Content
	message = Message{Content: m.Content, Author: "bot"}
	resp.Messages = append(resp.Messages, message)

	chatResp := ChatResponse{
		Candidates:     resp.Candidates,
		Filters:        resp.Filters,
		Messages:       resp.Messages,
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

func (c *ChatResponse) Reply(msg string) {
	message := Message{Content: msg, Author: "user"}
	c.Messages = append(c.Messages, message)
	msgConfig := MessageConfig{
		Prompt: MessagePrompt{
			Messages: c.Messages,
		},
	}
	apiKey, err := loadAPIKey(".env")
	if err != nil {
		return
	}
	msgResp, err := GenerateMessage(apiKey, msgConfig)
	if err != nil {
		return
	}
	resp := *msgResp
	botMsg := Message{Content: resp.Candidates[0].Content, Author: "bot"}
	c.Messages = append(c.Messages, botMsg)
}

func ChatPrompt(prompt string) (ChatResponse, error) {
	chatConfig := ChatConfig{Prompt: TextPrompt{Text: prompt}}
	resp, err := Chat(chatConfig)
	if err != nil {
		return ChatResponse{}, err
	}
	return resp, nil
}

func (c *Client) ChatPrompt(prompt string) (ResponseText, error) {
	response := ResponseText{}
	if c.config.authToken == "" {
		return response, fmt.Errorf("auth token not set")
	}
	endpoint := fmt.Sprintf("%s/%s:generateText?key=%s", API_BASE_URL, TEXT_MODEL, c.config.authToken)
	jsonMessagePrompt, err := json.Marshal(
		PromptConfig{
			Prompt: TextPrompt{
				Text: prompt,
			},
			MaxOutputTokens: 100,
			CandidateCount:  1,
			TopK:            1,
			TopP:            1,
		})
	if err != nil {
		return response, err
	}
	payload := string(jsonMessagePrompt)
	jsonPayload := []byte(payload)

	resp, err := makeRequest(endpoint, "POST", jsonPayload)
	if err != nil {
		return response, err
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (c *Client) Chat(config ChatConfig) (ChatResponse, error) {
	var msg string
	message := Message{Content: msg, Author: "user"}
	if len(config.Messages) == 0 && config.Prompt.Text != "" {
		msg = config.Prompt.Text
		message.Content = msg
		config.Messages = append(config.Messages, message)
	} else {
		msg = config.Messages[len(config.Messages)-1].Content
	}
	msgConfig := MessageConfig{
		Prompt: MessagePrompt{
			Messages: config.Messages,
		},
		Temperature:    config.Temperature,
		CandidateCount: config.CandidateCount,
		TopK:           config.TopK,
		TopP:           config.TopP,
	}
	msgResp, err := GenerateMessage(c.config.authToken, msgConfig)
	if err != nil {
		return ChatResponse{}, err
	}
	resp := *&msgResp
	m := resp.Candidates[0]
	message = Message{Content: m.Content, Author: "bot"}
	resp.Messages = append(resp.Messages, message)

	chatResp := ChatResponse{
		Candidates:     resp.Candidates,
		Filters:        resp.Filters,
		Messages:       resp.Messages,
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

func (c *Client) EmbedText(text string) (ResponseEmbed, error) {
    return EmbedText(text)
}
