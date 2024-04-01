package main

import (
	"context"
	"fmt"
)

func main() {
	apiKey := "..."
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	request := gpt3.CompletionRequest{
		Prompt: []string{"How many coffees should I drink per day?"},
	}
	resp, err := client.Completion(ctx, request)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Answer:\n %s\n", resp.Choices[0].Text)
	}

}
