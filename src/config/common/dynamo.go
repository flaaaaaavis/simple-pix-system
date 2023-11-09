package common

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

func GetItem(ctx context.Context, api DynamoDBAPI, params *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return api.GetItem(ctx, params)
}

func PutItem(ctx context.Context, api DynamoDBAPI, params *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return api.PutItem(ctx, params)
}
