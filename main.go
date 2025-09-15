package main

import (
	"context"
	"log"
	"log/slog"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{Logger: slog.New(slog.NewTextHandler(log.Writer(), &slog.HandlerOptions{Level: slog.LevelError}))})

	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, "fun-fact-tasks", worker.Options{})

	// Register the Workflow with the Worker
	w.RegisterWorkflow(FunFactWorkflow)

	// Register Activities
	w.RegisterActivity(PromptActivity)
	w.RegisterActivity(ReadInputActivity)
	w.RegisterActivity(GenAIActivity)

	// Start listening to the Task Queue
	err = w.Start()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

	// Start the workflow
	workflowOptions := client.StartWorkflowOptions{
		ID:        "fun-fact-workflow",
		TaskQueue: "fun-fact-tasks",
	}
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, FunFactWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	// Wait for workflow completion
	err = we.Get(context.Background(), nil)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow completed")
}
