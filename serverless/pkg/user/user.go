package user

import (
	"errors"
	"encoding/json"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

var(
	ErrorFailedToFetchRecord = "Failed to fetch Record"
	ErrorFailedToUnmarshalRecord = "Failed to unmarshal record"
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


func CreateUser(ctx context.Context,)(){}

func UpdateUser(ctx context.Context,)(){}

func DeleteUser(ctx context.Context,) error{}

