package dynamo_repo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"mentoria/src/config"
	model "mentoria/src/user/model/dynamo_model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ClientDynamo struct {
	Client    dynamodb.Client
	TableName string
}

func (d *ClientDynamo) CreateItem(ctx context.Context, user model.User) error {

	item, err := attributevalue.MarshalMap(user)

	if err != nil {
		log.Fatalf("Couldn't add item to table. Here's why: %v\n", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.TableName),
	}

	_, err = d.Client.PutItem(ctx, input)

	return nil

}

func (d *ClientDynamo) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
	var u *model.User

	input := dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: fmt.Sprint(user.ID),
			},
		},
		TableName: aws.String(d.TableName),
	}

	res, err := d.Client.GetItem(ctx, &input)
	if err != nil {
		fmt.Printf("erro ao obter item do dynamo: %s", err)
	}

	err = attributevalue.UnmarshalMap(res.Item, u)
	if err != nil {
		log.Fatalf("erro no unmarshal da response do dynamo: %s", err)
	}

	return u, nil
}

func NewDynamoClient(client dynamodb.Client, cfg *config.Config) ClientDynamo {
	return ClientDynamo{
		Client:    client,
		TableName: cfg.DynamoDBConfig.TableName,
	}
}
