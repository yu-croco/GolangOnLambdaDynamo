package model

type Timestamp struct {
	UserId  int      `dynamodbav:"userId"`
	startAt int      `dynamodbav:"startAt"`
	endAt   int      `dynamodbav:"endAt"`
}
