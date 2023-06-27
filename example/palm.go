package main

import (
	"fmt"
	"github.com/mr-destructive/palm"
)

func ChatReplyExample() {
	fmt.Println(palm.GetModel("chat-bison-001"))
	msg := "what is kafka?"
	chat, err := palm.ChatPrompt(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(chat.Last)
	chat.Reply("how this can be used in backend?")
	fmt.Println(chat.Last)
}
