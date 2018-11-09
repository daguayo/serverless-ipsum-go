package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/drhodes/golorem"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            lorem.Sentence(8,15),
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
