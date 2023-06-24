package palm

import (
	"testing"
)

func TestGenerateText(t *testing.T) {
	text, err := GenerateText("text-bison-001", PromptConfig{Prompt: TextPrompt{"hello world"}})
	if err != nil {
		t.Errorf("GenerateText failed: %v", err)
	}
	if text == "" {
		t.Error("GenerateText returned empty text")
	}
}

func TestGenerateMessage(t *testing.T) {
	message := MessagePrompt{
		Messages: []Message{
			Message{
				Content: "hello world",
			},
		},
	}
	msg, err := GenerateMessage(message, map[string]string{"model": "chat-bison-001"})
	if err != nil {
		t.Errorf("GenerateMessage failed: %v", err)
	}
	if msg == "" {
		t.Error("GenerateMessage returned empty message")
	}
}

func TestEmbedText(t *testing.T) {
	embed, err := EmbedText("Hello world!")
	if err != nil {
		t.Errorf("EmbedText failed: %v", err)
	}
	if len(embed.Embedding.Value) == 0 {
		t.Error("EmbedText returned no embeddings")
	}
}
