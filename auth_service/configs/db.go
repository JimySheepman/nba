package configs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ConnectDB() *dynamodb.DynamoDB {

	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(EnvRegion()),
		Credentials: credentials.NewStaticCredentials(EnvAccessKey(), EnvAccessSecretKey(), ""),
	})
	client := dynamodb.New(sess)

	_, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Connection to dynamoDB failed.", err)
	}
	fmt.Println("Connected to DynamoDB")
	return client
}

var DB = ConnectDB()
