package main

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// FunFactWorkflow is a Temporal workflow that prompts for user input and generates fun facts.
func FunFactWorkflow(ctx workflow.Context) error {
	// Activity options
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
	}
	ctxAO := workflow.WithActivityOptions(ctx, ao)

	for {
		// Prompt user
		err := workflow.ExecuteActivity(ctxAO, PromptActivity).Get(ctxAO, nil)
		if err != nil {
			return err
		}

		// Read user input
		var userInput string
		err = workflow.ExecuteActivity(ctxAO, ReadInputActivity).Get(ctxAO, &userInput)
		if err != nil {
			return err
		}

		if userInput == "exit" {
			break
		}

		// Generate fun fact
		prompt := "Tell me a fun fact about " + userInput
		var aiFact string
		err = workflow.ExecuteActivity(ctxAO, GenAIActivity, prompt).Get(ctxAO, &aiFact)
		if err != nil {
			return err
		}

		// Log the fun fact
		workflow.GetLogger(ctx).Info("Fun fact: " + aiFact)
	}

	return nil
}
