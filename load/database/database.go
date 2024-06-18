package db

import (
	"load/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	TABLE_NAME = "loads"
)

type DBClient interface {
	InsertLoad(load types.Load) error
	GetAllLoads() ([]types.Load, error)
}

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

func (d DynamoDbClient) GetAllLoads() ([]types.Load, error) {

	result, err := d.dbStore.Scan(&dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	})

	loads := []types.Load{}

	if err != nil {
		return loads, err
	}

	for _, row := range result.Items {
		var load types.Load

		err := dynamodbattribute.UnmarshalMap(row, &load)

		if err != nil {
			return []types.Load{}, err
		}

		loads = append(loads, load)
	}

	return loads, nil
}
