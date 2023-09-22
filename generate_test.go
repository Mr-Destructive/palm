package palm

import (
	"testing"
)

func TestGenerateText(t *testing.T) {
	text, err := GenerateText("text-bison-001", "", PromptConfig{Prompt: TextPrompt{"hello"}})
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
				Content: "hello",
			},
		},
	}
	msgConfig := MessageConfig{
		Prompt: message,
	}
	msg, err := GenerateMessage("", msgConfig)
	if err != nil {
		t.Errorf("GenerateMessage failed: %v", err)
	}
	if msg.Candidates[0].Content == "" {
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

func TestChat(t *testing.T) {
	chat, err := Chat(ChatConfig{
		Model: "chat-bison-001",
		Examples: []Example{
			Example{
				Input: Message{
					Content: "hello !",
				},
				Output: Message{
					Content: "hello world!",
				},
			},
		},
		Messages: []Message{
			Message{
				Content: "what are you!",
			},
		},
	})
	if err != nil {
		t.Errorf("Chat failed: %v", err)
	}
	if len(chat.Messages) == 0 {
		t.Error("Chat returned no messages")
	}
}

func TestChatReply(t *testing.T) {
	chat, err := Chat(ChatConfig{
		Model: "chat-bison-001",
		Examples: []Example{
			Example{
				Input: Message{
					Content: "hello world!",
				},
				Output: Message{
					Content: "what is that greeting used for?",
				},
			},
		},
		Messages: []Message{
			Message{
				Content: "what are you!",
			},
		},
	})
	if err != nil {
		t.Errorf("Chat failed: %v", err)
	}
	chat.Reply("what can you do for me!")
	if len(chat.Messages) == 0 {
		t.Error("Chat returned no messages")
	}
	last := chat.Candidates[len(chat.Candidates)-1].Content
	if chat.Last != last {
		t.Error("Chat returned wrong last message")
	}
}

func TestChatPrompt(t *testing.T) {
	chat, err := ChatPrompt("write a poem on a golang developer")
	if err != nil {
		t.Errorf("ChatPrompt failed: %v", err)
	}
	if len(chat.Messages) == 0 {
		t.Error("ChatPrompt returned no messages")
	}
}
