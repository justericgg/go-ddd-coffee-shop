package repository

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

const orderTableName = "Order"

type OrderRepository struct{}

func (o OrderRepository) GenerateID() (order.ID, error) {
	ddbClient, err := dynamo.GetClient()
	if err != nil {
		return order.ID{}, fmt.Errorf("get conn err in GenerateID() %w", err)
	}
	count, err := ddbClient.Count(orderTableName)
	if err != nil {
		return order.ID{}, fmt.Errorf("count err in GenerateID() %w", err)
	}

	id := order.NewID(count+1, time.Now())

	return id, nil
}

func (o OrderRepository) Save(order order.Order) {
	panic("implement me")
}
