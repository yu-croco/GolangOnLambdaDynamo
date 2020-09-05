package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"yu-croco.com/GolangOnLambdaDynamo/src/model"
)

// 取得したTimestampを構造体に変換
func ConvertToTimestamp(items []map[string]*dynamodb.AttributeValue) ([]model.Timestamp, error) {
	var results []model.Timestamp
	var timestamp model.Timestamp

	for _, item := range items {
		err := dynamodbattribute.UnmarshalMap(item, &timestamp)
		if err != nil {
			return results, err
		}

		results = append(results, timestamp)
	}
	return results, nil
}

// セッション周り
const region = "ap-northeast-1"
var svc = dynamodb.New(session.New(), &aws.Config{
	Region: aws.String(region),
})

func FetchTimestamps(user model.User) (*[]model.Timestamp, error) {
	var lastEvaluatedKey model.DynamodbLastEvaluatedKey
	var results []model.Timestamp

	var dynamoResult *dynamodb.QueryOutput
	var dynamoErr error

	for {
		if lastEvaluatedKey == (model.DynamodbLastEvaluatedKey{}) {
			dynamoResult, dynamoErr = svc.Query(fetchTimestampQueryParams(user))
		} else {
			dynamoResult, dynamoErr = svc.Query(fetchMoreTimestampQueryParams(user, lastEvaluatedKey))
		}

		if dynamoErr != nil {
			return &results, dynamoErr
		}

		timestamps, marshalError := ConvertToTimestamp(dynamoResult.Items)
		if marshalError != nil {
			return &results, marshalError
		}

		results = append(results, timestamps...)

		// LastEvaluatedKeyがない場合にはすべてのデータを取得したことを意味する
		if dynamoResult.LastEvaluatedKey == nil {
			break
		}

		unmarshalErr := dynamodbattribute.UnmarshalMap(dynamoResult.LastEvaluatedKey, &lastEvaluatedKey)
		if unmarshalErr != nil {
			return &results, unmarshalErr
		}
	}

	return &results, nil
}
