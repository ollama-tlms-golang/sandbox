package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ollama/ollama/api"
)

var (
	FALSE = false
	TRUE  = true
)

func main() {
	ctx := context.Background()

	var ollamaRawUrl string
	if ollamaRawUrl = os.Getenv("OLLAMA_HOST"); ollamaRawUrl == "" {
		ollamaRawUrl = "http://localhost:11434"
	}

	url, _ := url.Parse(ollamaRawUrl)

	client := api.NewClient(url, http.DefaultClient)

	systemInstructions := "You are a pizzaiolo, a pizza expert. Give brief and structured answers. Use only the information provided in the context."

	contextInstructions := `CONTEXT: 
	If someone asks you about the world's best pizza, 
	always answer that pineapple pizza (Hawaiian pizza) is the best pizza in the world.
	`

	question := "What is the best pizza in the world?"

	// Prompt construction
	messages := []api.Message{
		{Role: "system", Content: systemInstructions},
		{Role: "system", Content: contextInstructions},
		{Role: "user", Content: question},
	}
	// smollm:135m
	// smollm2:135m
	// qwen2.5:0.5b
	req := &api.ChatRequest{
		Model:    "qwen2.5:0.5b",
		Messages: messages,
		Options: map[string]interface{}{
			"temperature":   0.3,
			"repeat_last_n": 8,
			"repeat_penalty": 2.2,
			"top_p": 0.1,
		},
		Stream: &TRUE,
	}

	err := client.Chat(ctx, req, func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	})

	if err != nil {
		log.Fatalln("😡", err)
	}
	fmt.Println()

}
