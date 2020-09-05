package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
	"yu-croco.com/GolangOnLambdaDynamo/src/model"
)

const table = "timestampTable"

// 初回用のQuery
func fetchTimestampQueryParams(user model.User) *dynamodb.QueryInput {
	var limit int64 = 10000
	return &dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("#userId = :userId"),
		ExpressionAttributeNames: map[string]*string{
			"#userId": aws.String("userId"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {
				S: aws.String(string(user.UserId)),
			},
		},
		Limit: &limit,
	}
}

// 2回目以降用のQuery
// 前回のQueryで取得したlastEvaluatedKeyを使い、それ以降のデータを取得する
func fetchMoreTimestampQueryParams(user model.User, lastEvaluatedKey model.DynamodbLastEvaluatedKey) *dynamodb.QueryInput {
	var limit int64 = 10000
	startAt := strconv.Itoa(lastEvaluatedKey.StartAt)

	return &dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("#userId = :userId"),
		ExpressionAttributeNames: map[string]*string{
			"#userId": aws.String("userId"),
		},
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"startAt": {
				N: aws.String(startAt),
			},
			"userId": {
				S: aws.String(lastEvaluatedKey.UserId),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userId": {
				S: aws.String(string(user.UserId)),
			},
		},
		Limit: &limit,
	}
}