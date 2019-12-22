package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/infra/repository"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/usecase"
	"log"
	"time"
)

type Item struct {
	ProductID string `json:"ProductID"`
	Qty       int    `json:"Qty"`
	Price     int64  `json:"Price"`
}

type Event struct {
	ID         string    `json:"ID"`
	TableNo    string    `json:"TableNo"`
	Items      []Item    `json:"Items"`
	CratedTime time.Time `json:"CratedTime"`
}

func HandleRequest(ctx context.Context, e events.CloudWatchEvent) {
	log.Println("make coffee triggered", string(e.Detail))

	var event Event
	err := json.Unmarshal(e.Detail, &event)
	if err != nil {
		fmt.Printf("make coffee json parser error %v", err)
		return
	}
	log.Printf("event: %+v", event)
	items := make([]usecase.Item, 0)
	for _, i := range event.Items {
		items = append(items, usecase.Item{
			ProductID: i.ProductID,
			Qty:       i.Qty,
		})
	}
	cmd := usecase.MakeCoffeeCmd{
		TableNo: event.TableNo,
		OrderID: event.ID,
		Items:   items,
	}
	svc := usecase.NewMakeCoffeeSvc(repository.CoffeeRepository{}, repository.ProductRepository{})
	results, err := svc.MakeCoffee(cmd)
	if err != nil {
		log.Printf("make coffee error %v", err)
		return
	}

	fmt.Printf("coffee maked, OrderID: %s, results: %+v", cmd.OrderID, results)
}

func main() {
	lambda.Start(HandleRequest)
}
