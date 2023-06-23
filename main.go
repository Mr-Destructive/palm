package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnvFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

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
		m, err := GenerateMessage(message, map[string]string{"model": "chat-bison-001"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(m)
		model, err := GetModel(modelName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(model)
		ListModels()
	*/
	text, err := GenerateText(modelName, PromptConfig{Prompt: TextPrompt{"what is the meaning of life"}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
