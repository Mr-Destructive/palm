package palm

type TextPrompt struct {
	Text string `json:"text"`
}

type SafetySetting struct {
	StopSequence string `json:"stopSequence"`
}

type ChatConfig struct {
	Model          string     `json:"model;omitempty"`
	Context        string     `json:"context;omitempty"`
	Examples       []Example  `json:"examples,omitempty"`
	Messages       []Message  `json:"messages;omitempty"`
	Temperature    float64    `json:"temperature;omitempty"`
	CandidateCount int        `json:"candidateCount;omitempty"`
	TopP           float64    `json:"topP;omitempty"`
	TopK           int        `json:"topK;omitempty"`
	Prompt         TextPrompt `json:"prompt"`
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

type MessageConfig struct {
	Prompt         MessagePrompt `json:"prompt"`
	Temperature    float64       `json:"temperature"`
	CandidateCount int           `json:"candidateCount"`
	TopP           float64       `json:"topP"`
	TopK           int           `json:"topK"`
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

type ContentFilter struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type TextCompletion struct {
	Output           string           `json:"output"`
	SafetyRatings    []SafetyRating   `json:"safetyRatings"`
	CitationMetadata CitationMetadata `json:"citationMetadata,omitempty"`
}

type SafetyFeedback struct {
	Rating  SafetyRating  `json:"rating"`
	Setting SafetySetting `json:"setting"`
}

type SafetyRating struct {
	Category  string  `json:"category"`
	Threshold float64 `json:"threshold"`
}

type Embedding struct {
	Value []float64 `json:"value"`
}
