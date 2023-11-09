package config

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamoConnect *dynamodb.Client

func connection() error {
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Sprintf("Error loading aws default config: %s", err)
		return err
	}

	dynamoConnect = dynamodb.NewFromConfig(cfg)

	return nil
}

func NewDynamoDatabase() (*dynamodb.Client, error) {
	if dynamoConnect == nil {
		err := connection()
		if err != nil {
			fmt.Sprintf("Error when creating new dynamo connection: %s", err)
			return nil, err
		}

		return dynamoConnect, nil
	}

	return dynamoConnect, nil
}
