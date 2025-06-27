package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.QueryStringParameters["name"]
	if name == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Please provide a name parameter",
		}, nil
	}

	message := fmt.Sprintf("hii %s", name)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       message,
	}, nil
}

func main() {
	lambda.Start(handler)
}
