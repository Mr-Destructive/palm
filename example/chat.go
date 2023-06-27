package main

import (
	"github.com/mr-destructive/palm"
)

func ChatExample() {
	chatConfig := palm.ChatConfig{
		Model: "chat-bison-001",
		Examples: []palm.Example{
			palm.Example{
				Input: palm.Message{
					Author:  "Palm",
					Content: "hello world!",
				},
				Output: palm.Message{
					Author:  "Palm",
					Content: "hello world!",
				},
			},
		},
		Messages: []palm.Message{
			palm.Message{
				Content: "what are you!",
			},
		},
	}
	chat, err := palm.Chat(chatConfig)
	if err != nil {
		panic(err)
	}
	chat.Reply("what can you do for me!")
}

func ChatPromptExample() {
	chat, err := palm.ChatPrompt("how many continents are there?")
	if err != nil {
		panic(err)
	}
	chat.Reply("what lies in the amazon rainforest?")
}
