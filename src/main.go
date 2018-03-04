package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"helloworld/src/controller"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: controller.Hello(), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
