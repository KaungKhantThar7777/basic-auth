package database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "userTable"
)

type DynamodbClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDB() DynamodbClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamodbClient{
		databaseStore: db,
	}
}

func (client *DynamodbClient) DoesUserExists(username string) (bool, error) {
	result, err := client.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, err
	}

	return true, nil
}

func (client *DynamodbClient) InsertUser(user types.RegisterUser) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}

	_, error := client.databaseStore.PutItem(item)

	return error
}
