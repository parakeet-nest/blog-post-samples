package main

import (
	"fmt"
	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/llm"
)

var ollamaUrl = "http://localhost:11434"

// if working from a container
// use: "http://host.docker.internal:11434"

var systemContent = `You are an expert of the Star Trek franchise.`

var model = "qwen:0.5b"

func Question(userContent string) (string, error) {
	options := llm.Options{
		Temperature: 0.0, 
		RepeatLastN: 2,
		RepeatPenalty: 1.5,
	}

	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "user", Content: userContent},
		},
		Options: options,
	}

	answer, err := completion.Chat(ollamaUrl, query)

	fmt.Println("🙂 >", userContent)
	fmt.Println("🤖 >", answer.Message.Content)
	fmt.Println()

	if err != nil {
		return "", err
	}

	return answer.Message.Content, nil

}


func main() {
	userContent := "Who is James T Kirk and who is his best friend?"
	_, err := Question(userContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Jean-Luc Picard and who is his best friend?"
	_, err = Question(userContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Michael Burnham and who is his best friend?"
	_, err = Question(userContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Spock and who is his best friend?"
	_, err = Question(userContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	/*
		KeegOrg does not exist in the Star Trek franchise
		but the agent knows that Star Trek is the main topic
		so, the agent try to answer as if the character exists in the Star Trek franchise
	*/
	userContent = "Who is KeegOrg and who is his best friend?"
	_, err = Question(userContent)

	if err != nil {
		fmt.Println("😡:", err)
	}
}
