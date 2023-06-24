package main

import (
	"fmt"

	"github.com/mr-destructive/palm"
)

func EmbedTextExample() {
	embed, err := palm.EmbedText("Hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Println(embed)
}
