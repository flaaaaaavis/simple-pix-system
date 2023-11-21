package dynamo_repo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"mentoria/src/config/common"
	"mentoria/src/user/model/dynamo_model"
)

type DynamoClient struct {
	TableName string
	Client    common.DynamoDBAPI
}

const (
	PKFormat = "Pk#%s"
	SKFormat = "Sk#%s"
)

func (d *DynamoClient) GetUser(ctx context.Context, user *dynamo_model.User) (*dynamo_model.User, error) {
	dbUser := &dynamo_model.User{}

	selectedKeys := map[string]string{
		"PK": fmt.Sprintf(PKFormat, user.ID),
		"SK": fmt.Sprintf(SKFormat, user.ID),
	}

	key, err := attributevalue.MarshalMap(selectedKeys)
	if err != nil {
		log.Fatalf("Error on MarshalMap of the selected keys: %v\n", err)
		return nil, err
	}

	res := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(d.TableName),
	}

	item, err := d.Client.GetItem(ctx, res)
	if err != nil {
		log.Fatalf("Error getting item from dynamo: %v\n", err)
		return nil, err
	}

	err = attributevalue.UnmarshalMap(item.Item, dbUser)

	return dbUser, nil
}

func (d *DynamoClient) CreateItem(ctx context.Context, user *dynamo_model.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		log.Fatalf("Error converting user to Dynamo Type: %v\n", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.TableName),
	}

	_, err = d.Client.PutItem(ctx, input)
	if err != nil {
		log.Fatalf("Error sending item to dynamo: %v\n", err)
		return err
	}

	return nil
}
