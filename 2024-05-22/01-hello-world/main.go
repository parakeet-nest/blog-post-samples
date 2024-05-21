package main

import (
	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/llm"
	"fmt"
	"log"
)

func main() {
	ollamaUrl := "http://localhost:11434"
	// if working from a container
	//ollamaUrl := "http://host.docker.internal:11434"
	model := "tinydolphin"


	options := llm.Options{
		Temperature: 0.5, // default (0.8)
	}

	helloWorldQuery := llm.Query{
		Model: model,
		Prompt: `Generate a 'Hello, World!' program in Golang.`,
		Options: options,
	}
	
	_, err := completion.GenerateStream(ollamaUrl, helloWorldQuery,
		func(answer llm.Answer) error {
			fmt.Print(answer.Response)
			return nil
		})

	if err != nil {
		log.Fatal("😡:", err)
	}
}
