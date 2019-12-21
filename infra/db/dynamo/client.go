package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"sync"
)

const AwsRegion = "us-west-2"
const AwsCredProfile = "coffee_shop"

var (
	ddbClient *Client
	once      sync.Once
)

type Client struct {
	db *dynamodb.DynamoDB
}

func GetClient() (*Client, error) {
	var err error
	once.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(AwsRegion),
			Credentials: credentials.NewSharedCredentials("", AwsCredProfile),
		})
		if err == nil {
			ddbClient = &Client{db: dynamodb.New(sess)}
		}
	})
	return ddbClient, err
}

func (c *Client) Count(table string) (int64, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(table),
	}
	output, err := c.db.Scan(input)
	if err != nil {
		return 0, err
	}
	return *output.Count, nil
}
