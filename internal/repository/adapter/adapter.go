package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entiry interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
		logMode:    false,
	}
}

func (db *Database) Health() bool {
	// if tables are getting listed then health is ok
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	return db.connection.GetItemInput(input)
}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {

	conditionParsed, err := dynamodbattribute.MarshallMap(condition)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	}
	return db.connection.GetItemInput(input)
}

func (db *Database) CreateOrUpdate(entiry interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {

	entityParsed, err := dynamodbattribute.MarshallMap(entiry)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {

	conditionParsed, err := dynamodbattribute.MarshallMap(condition)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
