

package main

import (
	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/gear"
	"github.com/parakeet-nest/parakeet/llm"
	"github.com/parakeet-nest/parakeet/tools"
	"github.com/parakeet-nest/parakeet/wasm"

	"fmt"
	"log"
)

func main() {
	//ollamaUrl := "http://localhost:11434"
	// if working from a container
	ollamaUrl := "http://host.docker.internal:11434"
	model := "mistral:7b"

	toolsList := []llm.Tool{
		{
			Type: "function",
			Function: llm.Function{
				Name:        "hello",
				Description: "Say hello to a given person with his name",
				Parameters: llm.Parameters{
					Type: "object",
					Properties: map[string]llm.Property{
						"name": {
							Type:        "string",
							Description: "The name of the person",
						},
					},
					Required: []string{"name"},
				},
			},
		},
		{
			Type: "function",
			Function: llm.Function{
				Name:        "hey",
				Description: "Say hey to a given person with his name",
				Parameters: llm.Parameters{
					Type: "object",
					Properties: map[string]llm.Property{
						"name": {
							Type:        "string",
							Description: "The name of the person",
						},
					},
					Required: []string{"name"},
				},
			},
		},
	}

	toolsContent, err := tools.GenerateContent(toolsList)
	if err != nil {
		log.Fatal("😡:", err)
	}

	userContent := tools.GenerateInstructions(`say "hey" to Sam`)


	options := llm.Options{
		Temperature:   0.0,
		RepeatLastN:   2,
		RepeatPenalty: 2.0,
	}

	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: toolsContent},
			{Role: "user", Content: userContent},
		},
		Options: options,
		Format:  "json",
		Raw:     true,
	}

	answer, err := completion.Chat(ollamaUrl, query)
	if err != nil {
		log.Fatal("😡:", err)
	}

	jsonRes, err := gear.JSONParse(answer.Message.Content)

	if err != nil {
		log.Fatal("😡:", err)
	}
	functionName := jsonRes["name"].(string)
	name := jsonRes["arguments"].(map[string]interface{})["name"].(string)

	fmt.Println("Calling", functionName, "with", name)

	wasmPlugin, _ := wasm.NewPlugin("./plugin/target/wasm32-unknown-unknown/debug/greetings.wasm", nil)

	// call the function of the wasm plugin
	res, _ := wasmPlugin.Call(functionName, []byte(name))

	// display the result
	fmt.Println(string(res))


}
