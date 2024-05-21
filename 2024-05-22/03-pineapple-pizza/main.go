package main

import (
	"fmt"
	"log"

	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/llm"
)

func main() {
	ollamaUrl := "http://localhost:11434"
	// if working from a container
	//ollamaUrl := "http://host.docker.internal:11434"
	model := "tinydolphin"

	systemContent := `You are an expert on the pizza topic? 
	Based on the given examples and your knowledge, 
	give the more accurate answer.`
	contextContent := `
	QUESTION: What is the best pizza of the world?
	ANSWER: The pineapple pizza is the best pizza of the world for the following reasons:
	- Sweet and Savory Fusion.
	- Textural Delight.
	- Balanced Flavors.
	- Tropical Twist.
	`

	options := llm.Options{
		Temperature: 0.5, // default (0.8)
	}

	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "system", Content: contextContent},
			{Role: "user", Content: `What is the best pizza of the world?`},
		},
		Options: options,
	}

	_, err := completion.ChatStream(ollamaUrl, query,
		func(answer llm.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	if err != nil {
		log.Fatal("😡:", err)
	}

}