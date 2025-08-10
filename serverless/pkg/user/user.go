package user

import (
	"context"
	"encoding/json"
	"errors"
	"serverless/pkg/validators"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var(
	ErrorFailedToFetchRecord = "Failed to fetch Record"
	ErrorFailedToUnmarshalRecord = "Failed to unmarshal record"
	ErrorInvalidUserData = "Inavalid user data"
	ErrorInvalidEmail = "Invalid email"
	ErrorCouldNotMarshalItem = "Could not marshal item"
	ErrorCouldNotDeleteItem = "Could not delete item"
	ErrorCouldNotDynamoPutItem = "Could not dynamo put item"
	ErrorUserAlreadyExists = "user.User already exists"
	ErrorUserDoesNotExist = "user does not exist"
)

type User struct{
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func FetchUser(ctx context.Context, email string, tableName string, dynaClient *dynamodb.Client) (*User, error) {
    input := &dynamodb.GetItemInput{
        Key: map[string]types.AttributeValue{
            "email": &types.AttributeValueMemberS{
                Value: email,
            },
        },
        TableName: aws.String(tableName),
    }

    res, err := dynaClient.GetItem(ctx, input)
    if err != nil {
        return nil, errors.New(ErrorFailedToFetchRecord)
    }

    if res.Item == nil {
        return nil, errors.New("user not found")
    }

    item := new(User)
    err = attributevalue.UnmarshalMap(res.Item, item)
    if err != nil {
        return nil, errors.New(ErrorFailedToUnmarshalRecord)
    }

    return item, nil
}


func FetchUsers(ctx context.Context, tableName string, dynaClient *dynamodb.Client) (*[]User, error) {
    input := &dynamodb.ScanInput{
        TableName: aws.String(tableName),
    }

    res, err := dynaClient.Scan(ctx, input)
    if err != nil {
        return nil, errors.New(ErrorFailedToFetchRecord)
    }

    var users []User
    err = attributevalue.UnmarshalListOfMaps(res.Items, &users)
    if err != nil {
        return nil, errors.New(ErrorFailedToUnmarshalRecord)
    }

    return &users, nil
}


func CreateUser(ctx context.Context,req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client)(*User, error){
	var u User
	if err:= json.Unmarshal([]byte(req.Body),&u); err!=nil{
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email){
		return nil, errors.New(ErrorInvalidEmail)
	}
	currentUser,_:=FetchUser(ctx,u.Email,tableName,dynaClient)
	if currentUser!=nil && len(currentUser.Email)!=0{
		return nil, errors.New(ErrorUserAlreadyExists)
	}
	av,err := attributevalue.MarshalMap(u)
	if err!=nil{
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}
	input:= &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(tableName),
	}
	_,err= dynaClient.PutItem(ctx, input)
	if err!=nil{
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil;
}

func UpdateUser(ctx context.Context,)(){}

func DeleteUser(ctx context.Context,) error{}

