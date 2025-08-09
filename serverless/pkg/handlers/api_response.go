package handlers

import(
	 "encoding/json"
	 "github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }

    res := events.APIGatewayProxyResponse{
        StatusCode: status,
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
        Body: string(jsonBody),
    }

    return &res, nil
}
