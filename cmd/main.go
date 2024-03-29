package main

import (
	"dynamodb-crud/internal/repository/adapter"
	"dynamodb-crud/internal/repository/instance"
	"dynamodb-crud/internal/routes"
	"dynamodb-crud/utils/logger"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	Config "github.com/borowiak-m/dynamodb-crud/config"
)

func main() {
	configs := Config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("waiting for the service to start...", nil)

	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migration:.....", err)
		}
	}

	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("Service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTables(&dynamodb.ListTablesInput()))

	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables not found", nil)
		}

		for _, tableName := range response.TanleNames {
			logger.INFO("Table found", *tableName)
		}
	}
	return err
}
