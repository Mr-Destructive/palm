package main

import (
	"fmt"
	"github.com/mr-destructive/palm"
)

func main() {
	modelName := "text-bison-001"
	/*
		message := MessagePrompt{
			Context: "what is the meaning of life",
			Examples: []Example{
				Example{
					Input: Message{
						Author:  "Palm",
						Content: "what is the meaning of life",
					},
					Output: Message{
						Author:  "Palm",
						Content: "what is the meaning of life",
					},
				},
			},
			Messages: []Message{
				Message{
					Author:  "Palm",
					Content: "what is the meaning of life",
				},
			},
		}
	*/
	message := palm.MessagePrompt{
		Messages: []palm.Message{
			palm.Message{
				Content: "what is the meaning of life",
			},
		},
	}
	m, err := palm.GenerateMessage(message, map[string]string{"model": "chat-bison-001"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
	model, err := palm.GetModel(modelName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(model)
	palm.ListModels()
	text, err := palm.GenerateText(modelName, palm.PromptConfig{Prompt: palm.TextPrompt{"what is the meaning of life"}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
