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
	ListModels()
	text, err := GenerateText("what is the meaning of life", map[string]string{"model": modelName})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
