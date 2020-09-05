package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"yu-croco.com/GolangOnLambdaDynamo/src/dynamodb"
	"yu-croco.com/GolangOnLambdaDynamo/src/model"
)

func Handler() error {
	user := model.User{
		UserId: 1,
		UserName: "Taro",
	}

	timestamps, fetchErr := dynamodb.FetchTimestamps(user)
	if fetchErr != nil {
		return fetchErr
	}

	fmt.Println(timestamps)

	return nil
}

func main() {
	lambda.Start(Handler)
}
