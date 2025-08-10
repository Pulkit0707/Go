package main

import (
    "context"
    "os"

    "serverless/pkg/handlers"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
    dynaClient *dynamodb.Client
    tableName  = "LambdaInGoUser"
)

func main() {
    ctx := context.Background()

    // Load AWS configuration from environment/credentials
    cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(os.Getenv("AWS_REGION")))
    if err != nil {
        panic("Unable to load AWS config: " + err.Error())
    }

    // Initialize DynamoDB client
    dynaClient = dynamodb.NewFromConfig(cfg)

    // Start Lambda
    lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    switch req.HTTPMethod {
    case "GET":
        return handlers.GetUser(ctx, req, tableName, dynaClient)
    case "POST":
        return handlers.CreateUser(ctx, req, tableName, dynaClient)
    case "PUT":
        return handlers.UpdateUser(ctx, req, tableName, dynaClient)
    case "DELETE":
        return handlers.DeleteUser(ctx, req, tableName, dynaClient)
    default:
        return handlers.UnhandledMethod()
    }
}
