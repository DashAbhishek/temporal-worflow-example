package main

import (
	"context"
	"fmt"
	"log"
	"play/app"
	"play/app/activities"
	"play/app/workflows"

	"go.temporal.io/sdk/client"
)

func main() {

	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "joke-workflow",
		TaskQueue: app.JokeTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.JokeWorkFlow, "ghost")
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var jokes []activities.Joke
	err = we.Get(context.Background(), &jokes)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	printResults(jokes, we.GetID(), we.GetRunID())
}

func printResults(jokes []activities.Joke, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", jokes)

}
