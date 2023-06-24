package palm

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

type ContentFilter struct {
	Reason  BlockedReason `json:"reason"`
	Message string        `json:"message"`
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

type HarmCategory int

const (
	HARM_CATEGORY_UNSPECIFIED HarmCategory = iota
	HARM_CATEGORY_DEROGATORY
	HARM_CATEGORY_TOXICITY
	HARM_CATEGORY_VIOLENCE
	HARM_CATEGORY_SEXUAL
	HARM_CATEGORY_MEDICAL
	HARM_CATEGORY_DANGEROUS
)

func (c HarmCategory) String() string {
	switch c {
	case HARM_CATEGORY_UNSPECIFIED:
		return "HARM_CATEGORY_UNSPECIFIED"
	case HARM_CATEGORY_DEROGATORY:
		return "HARM_CATEGORY_DEROGATORY"
	case HARM_CATEGORY_TOXICITY:
		return "HARM_CATEGORY_TOXICITY"
	case HARM_CATEGORY_VIOLENCE:
		return "HARM_CATEGORY_VIOLENCE"
	case HARM_CATEGORY_SEXUAL:
		return "HARM_CATEGORY_SEXUAL"
	case HARM_CATEGORY_MEDICAL:
		return "HARM_CATEGORY_MEDICAL"
	case HARM_CATEGORY_DANGEROUS:
		return "HARM_CATEGORY_DANGEROUS"
	}
	return ""
}

type HarmBlockThreshold int

const (
	HARM_BLOCK_THRESHOLD_UNSPECIFIED HarmBlockThreshold = iota
	BLOCK_LOW_AND_ABOVE
	BLOCK_MEDIUM_AND_ABOVE
	BLOCK_ONLY_HIGH
	BLOCK_NONE
)

func (t HarmBlockThreshold) String() string {
	switch t {
	case HARM_BLOCK_THRESHOLD_UNSPECIFIED:
		return "HARM_BLOCK_THRESHOLD_UNSPECIFIED"
	case BLOCK_LOW_AND_ABOVE:
		return "BLOCK_LOW_AND_ABOVE"
	case BLOCK_MEDIUM_AND_ABOVE:
		return "BLOCK_MEDIUM_AND_ABOVE"
	case BLOCK_ONLY_HIGH:
		return "BLOCK_ONLY_HIGH"
	case BLOCK_NONE:
		return "BLOCK_NONE"
	}
	return ""

}

type Embedding struct {
    Value float64 `json:"value"`
}
