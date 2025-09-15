package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

// PromptActivity prompts the user for input.
func PromptActivity(ctx context.Context) error {
	fmt.Print("Enter a word to discover a fun fact, or type 'exit' to quit: ")
	return nil
}

// ReadInputActivity reads user input from stdin.
func ReadInputActivity(ctx context.Context) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}
	return "", scanner.Err()
}

// GenAIActivity is a Temporal activity that uses Genkit to generate AI text.
func GenAIActivity(ctx context.Context, prompt string) (string, error) {
	o := &ollama.Ollama{ServerAddress: "http://127.0.0.1:11434"}
	g := genkit.Init(context.Background(), genkit.WithPlugins(o))

	modelDef := ollama.ModelDefinition{
		Name: "gemma3:1b",
		Type: "generate", // "chat" or "generate"
	}

	// Define the model and get the model instance
	model := o.DefineModel(g, modelDef, nil)
	if model == nil {
		return "", fmt.Errorf("failed to define model: gemma3:1b")
	}

	resp, err := genkit.Generate(ctx, g, ai.WithModel(model), ai.WithPrompt(prompt))
	if err != nil {
		return "", err
	}
	fmt.Println("AI response:", resp.Text())
	return resp.Text(), nil
}
