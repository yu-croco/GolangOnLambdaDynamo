package model

// dynamoDBから返されたLastEvaluatedKeyを格納する
type DynamodbLastEvaluatedKey struct {
	UserId   string `dynamodbav:"userId"`
	StartAt int    `dynamodbav:"startAt"`
}