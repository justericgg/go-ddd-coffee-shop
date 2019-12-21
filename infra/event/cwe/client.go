package cwe

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/justericgg/go-ddd-coffee-shop/ddd"
	"log"
	"sync"
)

const (
	AwsRegion      = "us-west-2"
	AwsCredProfile = "coffee_shop"
)

var (
	cweClient *Client
	once      sync.Once
)

type Client struct {
	cwe *cloudwatchevents.CloudWatchEvents
}

func GetClient() *Client {
	once.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(AwsRegion),
			Credentials: credentials.NewSharedCredentials("", AwsCredProfile),
		})
		if err == nil {
			cweClient = &Client{cwe: cloudwatchevents.New(sess)}
		}
	})

	return cweClient
}

func (c *Client) Publish(events []ddd.DomainEvent) error {
	for _, event := range events {
		err := c.putEvent(event)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) putEvent(event ddd.DomainEvent) error {
	msg, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	result, err := c.cwe.PutEvents(&cloudwatchevents.PutEventsInput{
		Entries: []*cloudwatchevents.PutEventsRequestEntry{
			{
				Detail:     aws.String(string(msg)),
				DetailType: aws.String("OrderCreated"),
				Resources: []*string{
					aws.String("arn:aws:events:us-west-2:378652145250:rule/OrderCreated"),
				},
				Source: aws.String("com.event.justericgg"),
			},
		},
	})
	if err != nil {
		return err
	}

	log.Println("ingested event:", result.Entries, string(msg))

	return nil
}
