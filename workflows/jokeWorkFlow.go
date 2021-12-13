package workflows

import (
	"play/app/activities"
	"time"

	"go.temporal.io/sdk/workflow"
)

func JokeWorkFlow(ctx workflow.Context, searchval string) ([]activities.Joke, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result []activities.Joke
	err := workflow.ExecuteActivity(ctx, activities.GetJoke, searchval).Get(ctx, &result)
	return result, err
}
