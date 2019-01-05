package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var dynamoDbSession = session.Must(session.NewSession())
var cfg = aws.Config{
	Endpoint: aws.String("http://localhost:9000"),
	Region:   aws.String("eu-west-1"),
}
var DynamoDbClient = dynamo.New(dynamoDbSession, &cfg)

func GetDynamodbTable(tableName string) (*dynamo.Table, error) {
	if tableName == "" {
		return nil, log.Error("you must supply a table name")
	}
	table := DynamoDbClient.Table(tableName)
	return &table, nil
}
