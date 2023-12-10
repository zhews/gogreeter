package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zhews/gogreeter/greeter"
)

func HandleRequest(ctx context.Context, event events.APIGatewayV2HTTPRequest) ([]byte, error) {
	name := event.QueryStringParameters["name"]
	greeting := greeter.Greet(name)
	output, err := json.Marshal(greeting)
	if err != nil {
		return nil, fmt.Errorf("could not marshal greeting %w", err)
	}
	return output, nil
}

func main() {
	lambda.Start(HandleRequest)
}
