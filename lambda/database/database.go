package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	TABLE_NAME = "userTable"
)

type UserStore interface {
	DoesUserExists(username string) (bool, error)
	InsertUser(user types.User) error
	GetUser(username string) (types.User, error)
}
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

func (client DynamodbClient) DoesUserExists(username string) (bool, error) {
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

func (client DynamodbClient) InsertUser(user types.User) error {
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.PasswordHash),
			},
		},
	}

	_, error := client.databaseStore.PutItem(item)

	return error
}

func (client DynamodbClient) GetUser(username string) (types.User, error) {
	var user types.User

	result, err := client.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return user, err
	}

	if result.Item == nil {
		return user, fmt.Errorf("user not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	return user, err
}
