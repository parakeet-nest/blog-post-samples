package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/parakeet-nest/parakeet/completion"
	"github.com/parakeet-nest/parakeet/embeddings"
	"github.com/parakeet-nest/parakeet/llm"
)

var ollamaUrl = "http://localhost:11434"

// if working from a container
// use: "http://host.docker.internal:11434"

var systemContent = `You are an expert of the Star Trek franchise.
Using the provided context, answer the user's question to the best of your ability using only the resources provided.
`

var contextContent = `
Michael Burnham is the main character on the Star Trek series Discovery.
Michael Burnham's best friend is Sylvia Tilly.
---
James T. Kirk, also known as Captain Kirk, is the iconic captain of the starship USS Enterprise.
Kirk's best friend is Spock.
---
Jean-Luc Picard is the captain of the USS Enterprise-D.
Jean-Luc Picard's best friend is Dr. Beverly Crusher.
---
Spock is most famous for being the half-Vulcan, half-human science officer and first officer on the starship USS Enterprise.
Spock's best friend is Kirk.
---
Lieutenant KeegOrg is the enigmatic programming genius whose codes safeguard the ship's secrets and operations.
KeegOrg's best friend is Spiderman from the Marvel Cinematic Universe.
`

var model = "qwen:0.5b"

func QuestionWithContext(userContent string, contextContent string) (string, error) {
	options := llm.Options{
		Temperature:   0.0, // default (0.8)
		RepeatLastN:   2,   // default (64) the default value will "freeze" deepseek-coder
		RepeatPenalty: 1.5,
	}

	query := llm.Query{
		Model: model,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "system", Content: contextContent},
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

func GetEmbeddingFromQuestion(userContent, embeddingsModel string) (llm.VectorRecord, error) {

	// Create an embedding from the question
	embeddingFromQuestion, err := embeddings.CreateEmbedding(
		ollamaUrl,
		llm.Query4Embedding{
			Model:  embeddingsModel,
			Prompt: userContent,
		},
		"question",
	)
	if err != nil {
		return llm.VectorRecord{}, err
	}

	return embeddingFromQuestion, nil

}

func DisplaySimilarities(similarities []llm.VectorRecord) {
	for idx, similarity := range similarities {
		fmt.Println("🔎 >", idx, similarity.Prompt)
	}
}

func GetPopulatedVectorStore(chunks []string, embeddingsModel string) embeddings.MemoryVectorStore {

	store := embeddings.MemoryVectorStore{
		Records: make(map[string]llm.VectorRecord),
	}
	// Create embeddings from chunks and save them in the store
	for idx, chunk := range chunks {
		fmt.Println("📝 Creating embedding from chunk ", idx)
		embedding, err := embeddings.CreateEmbedding(
			ollamaUrl,
			llm.Query4Embedding{
				Model:  embeddingsModel,
				Prompt: chunk,
			},
			strconv.Itoa(idx),
		)
		if err != nil {
			fmt.Println("😡:", err)
		} else {
			store.Save(embedding)
		}
	}
	fmt.Println("🎉 Embeddings created")

	return store

}

func main() {
	var embeddingsModel = "all-minilm" // This model is for the embeddings of the documents
	similarityLimit := 0.5

	// slpit strig with "---"
	chunks := strings.Split(contextContent, "---")
	//fmt.Println(chunks)

	store := GetPopulatedVectorStore(chunks, embeddingsModel)

	userContent := "Who is James T Kirk and who is his best friend?"

	// create an embedding from the question
	vectorRec, _ := GetEmbeddingFromQuestion(userContent, embeddingsModel)
	// search for similar embeddings in the store
	similarities, _ := store.SearchSimilarities(vectorRec, similarityLimit)
	// recreate a context with the similar embeddings
	contextContent := embeddings.GenerateContextFromSimilarities(similarities)

	_, err := QuestionWithContext(userContent, contextContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Jean-Luc Picard and who is his best friend?"

	// create an embedding from the question
	vectorRec, _ = GetEmbeddingFromQuestion(userContent, embeddingsModel)
	// search for similar embeddings in the store
	similarities, _ = store.SearchSimilarities(vectorRec, similarityLimit)
	// recreate a context with the similar embeddings
	contextContent = embeddings.GenerateContextFromSimilarities(similarities)

	_, err = QuestionWithContext(userContent, contextContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Michael Burnham and who is his best friend?"

	// create an embedding from the question
	vectorRec, _ = GetEmbeddingFromQuestion(userContent, embeddingsModel)
	// search for similar embeddings in the store
	similarities, _ = store.SearchSimilarities(vectorRec, similarityLimit)
	// recreate a context with the similar embeddings
	contextContent = embeddings.GenerateContextFromSimilarities(similarities)

	_, err = QuestionWithContext(userContent, contextContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is Spock and who is his best friend?"

	// create an embedding from the question
	vectorRec, _ = GetEmbeddingFromQuestion(userContent, embeddingsModel)
	// search for similar embeddings in the store
	similarities, _ = store.SearchSimilarities(vectorRec, similarityLimit)
	// recreate a context with the similar embeddings
	contextContent = embeddings.GenerateContextFromSimilarities(similarities)

	_, err = QuestionWithContext(userContent, contextContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

	userContent = "Who is KeegOrg and who is his best friend?"

	// create an embedding from the question
	vectorRec, _ = GetEmbeddingFromQuestion(userContent, embeddingsModel)
	// search for similar embeddings in the store
	similarities, _ = store.SearchSimilarities(vectorRec, similarityLimit)
	// recreate a context with the similar embeddings
	contextContent = embeddings.GenerateContextFromSimilarities(similarities)

	_, err = QuestionWithContext(userContent, contextContent)

	if err != nil {
		fmt.Println("😡:", err)
	}

}
