package db

import (
	"new-load/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	TABLE_NAME = "loads"
)

type DynamoDbClient struct {
	dbStore *dynamodb.DynamoDB
}

func NewDynamoDb() DynamoDbClient {

	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDbClient{
		dbStore: db,
	}
}

func (d DynamoDbClient) InsertLoad(load types.Load) error {

	item, err := dynamodbattribute.MarshalMap(load)

	if err != nil {
		return err
	}

	_, err = d.dbStore.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      item,
	})

	if err != nil {
		return err
	}
	return nil
}
