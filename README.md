# Temporal Fun Fact Generator

A Temporal workflow application that creates a human-in-the-loop interaction to generate fun facts about user-provided words using Google's Genkit framework for AI integration.

## Features

- Interactive CLI workflow that prompts users for input
- Generates fun facts using Ollama and Gemma3:1b model
- Loops until user types "exit"
- Uses Temporal for workflow orchestration
- Structured logging with slog
- Leverages Genkit's advanced AI capabilities:
  - Type-safe AI flows with Go structs and JSON schema validation
  - Unified model interface supporting Google AI, Vertex AI, OpenAI, Ollama, and more
  - Tool calling, RAG, and multimodal support
  - Rich local development tools with a standalone CLI binary and Developer UI
  - AI coding assistant integration

## Prerequisites

- Go 1.21 or later
- Temporal server running locally
- Ollama with Gemma3:1b model installed

## Setup

1. **Install Temporal CLI and start server:**
   ```bash
   # Install Temporal CLI
   curl -sSf https://temporal.download/cli.sh | sh

   # Start Temporal server
   temporal server start-dev
   ```

2. **Install Ollama and model:**
   ```bash
   # Install Ollama
   curl -fsSL https://ollama.ai/install.sh | sh

   # Pull the Gemma3:1b model
   ollama pull gemma3:1b
   ```

3. **Install Genkit CLI:**
   ```bash
   curl -sL cli.genkit.dev | bash
   ```

4. **Clone and setup the project:**
   ```bash
   git clone https://github.com/AlexSKuznetsov/temporal-genkit.git
   cd temp-genai
   go mod tidy
   ```

## Running the Application

1. Ensure Temporal server is running in one terminal:
   ```bash
   temporal server start-dev
   ```

2. In another terminal, start Ollama:
   ```bash
   ollama serve
   ```

3. Run the application:
   ```bash
   genkit start -- go run .
   ```

The application will start the workflow and begin prompting for user input.

## How It Works

1. **Workflow Execution:**
   - The `FunFactWorkflow` starts and enters a loop
   - Executes `PromptActivity` to display the prompt
   - Executes `ReadInputActivity` to read user input
   - If input is "exit", workflow completes
   - Otherwise, executes `GenAIActivity` to generate a fun fact
   - Logs the fun fact and loops back

2. **Activities:**
   - `PromptActivity`: Displays the user prompt
   - `ReadInputActivity`: Reads input from stdin
   - `GenAIActivity`: Uses Ollama to generate AI responses

3. **Logging:**
   - Uses slog with info level to suppress debug logs
   - Fun facts are logged at info level

## Project Structure

```
.
├── main.go          # Worker setup and main function
├── workflow.go      # FunFactWorkflow definition
├── activity.go      # Activity definitions
├── go.mod           # Go module file
├── go.sum           # Go dependencies
└── README.md        # This file
```

## Configuration

- **Temporal Task Queue:** `fun-fact-tasks`
- **Workflow ID:** `fun-fact-workflow`
- **Ollama Server:** `http://127.0.0.1:11434`
- **Model:** `gemma3:1b`

## Usage Example

```
$ go run .
Started workflow WorkflowID fun-fact-workflow RunID ...
Enter a word to discover a fun fact, or type 'exit' to quit: apple
INFO Fun fact: Did you know that the average apple contains about 5-10 seeds, and if you plant them, you might grow a tree that produces fruit different from the parent apple?
Enter a word to discover a fun fact, or type 'exit' to quit: ocean
INFO Fun fact: The ocean covers about 71% of Earth's surface and contains 97% of Earth's water.
Enter a word to discover a fun fact, or type 'exit' to quit: exit
Workflow completed
```

## UI Links

- [Temporal UI](http://127.0.0.1:8233/) - Monitor workflows and activities
- [Genkit UI](http://localhost:4000/) - View AI model interactions and debugging

## Troubleshooting

- **Temporal server not running:** Ensure `temporal server start-dev` is executed
- **Ollama not running:** Run `ollama serve` in a separate terminal
- **Model not found:** Run `ollama pull gemma3:1b`
- **Port conflicts:** Ensure ports 7233 (Temporal) and 11434 (Ollama) are available

## About Genkit

This project uses [Genkit](https://genkit.dev/?lang=go), Google's open-source framework for building AI-powered apps. Genkit provides a unified interface for integrating various AI models and services, making it easy to add AI capabilities to applications.

Key features used in this project:
- Model abstraction for easy switching between AI providers
- Plugin system for Ollama integration
- Structured prompts and response handling

For more information and documentation, visit [Genkit for Go](https://genkit.dev/?lang=go).

## Dependencies

- `go.temporal.io/sdk`: Temporal Go SDK
- `github.com/firebase/genkit/go`: Genkit for AI integration
- `github.com/firebase/genkit/go/plugins/ollama`: Ollama plugin for Genkit
