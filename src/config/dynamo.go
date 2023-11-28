package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func CreateTables(dynamoConf DynamoConfig) (*dynamodb.Client, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
		opts.Region = dynamoConf.Region
		return nil
	})
	if err != nil {
		panic(err)
		return nil, err
	}

	svc := dynamodb.NewFromConfig(cfg, func(db *dynamodb.Options) {
		db.EndpointResolver = dynamodb.EndpointResolverFromURL("http://localhost:8000")
	})

	_, err = svc.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(dynamoConf.TableName),
		BillingMode: types.BillingModePayPerRequest,
	})

	return svc, nil
}
