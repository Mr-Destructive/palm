package main

import (
	"fmt"

	"github.com/mr-destructive/palm"
)

func main() {
	embed, err := palm.EmbedText("Hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Println(embed)
}
