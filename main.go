package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"os"
)

func main() {

	fmt.Println("hello world")

	svc := dynamoDb()

	res, err := svc.ListTablesRequest(&dynamodb.ListTablesInput{}).Send()
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.TableNames, len(res.TableNames))
}

func dynamodbConfig(baseConfig aws.Config) aws.Config {

	baseConfig.EndpointResolver = aws.ResolveWithEndpointURL("http://localhost:8000")

	//baseConfig.LogLevel = aws.LogDebugWithHTTPBody
	return baseConfig
}

func dynamoDb() *dynamodb.DynamoDB {
	d := dynamodb.New(dynamodbConfig(loadBasicAwsConfig()))
	return d
}

func loadBasicAwsConfig() aws.Config {

	os.Setenv("AWS_REGION", "us-east-1")

	cfg, err := external.LoadDefaultAWSConfig()
	cfg.LogLevel = aws.LogDebugWithHTTPBody
	if err != nil {
		log.Fatal(err)
	}

	return cfg.Copy()
}
