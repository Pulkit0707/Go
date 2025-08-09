package handlers

import (
	"net/http"
	"serverless/pkg/user"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct{
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*events.APIGatewayProxyResponse, error) {
    email := req.QueryStringParameters["email"]

    if len(email) > 0 {
        res, err := user.FetchUser(ctx, email, tableName, dynaClient)
        if err != nil {
            return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
        }
        return apiResponse(http.StatusOK, res)
    }

    res, err := user.FetchUsers(ctx, tableName, dynaClient)
    if err != nil {
        return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
    }
    return apiResponse(http.StatusOK, res)
}

func UpdateUser(ctx context.Context,req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodb.Client)(*events.APIGatewayProxyResponse, error){
	res,err:=user.UpdateUser(ctx, req, tableName, dynaClient)
	if err!=nil{
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, res)
}

func CreateUser(ctx context.Context,req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodb.Client)(*events.APIGatewayProxyResponse, error){
	res,err:=user.CreateUser(ctx, req, tableName, dynaClient)
	if err!=nil{
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, res)
}

func DeleteUser(ctx context.Context,req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodb.Client)(*events.APIGatewayProxyResponse, error){
	err:=user.DeleteUser(ctx, req, tableName, dynaClient)
	if err!=nil{
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod()(*events.APIGatewayProxyResponse,error){
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}