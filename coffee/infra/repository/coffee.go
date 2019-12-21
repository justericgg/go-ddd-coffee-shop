package repository

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/coffee"
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

const TableName = "Coffee"

type Schema struct {
	CoffeeID    string
	TableNo     string
	ProductName string
	CreatedAt   time.Time
}

type OrderRepository struct{}

func (o OrderRepository) GenerateID() (order.ID, error) {
	ddbClient, err := dynamo.GetClient()
	if err != nil {
		return order.ID{}, fmt.Errorf("get conn err in GenerateID() %w", err)
	}
	count, err := ddbClient.Count(TableName)
	if err != nil {
		return order.ID{}, fmt.Errorf("count err in GenerateID() %w", err)
	}

	id := order.NewID(count+1, time.Now())

	return id, nil
}

func (o OrderRepository) Save(cof *coffee.Coffee) error {
	ddbClient, err := dynamo.GetClient()
	if err != nil {
		return fmt.Errorf("get conn err in Save() %w", err)
	}

	input := &Schema{
		CoffeeID:    cof.ID().String(),
		TableNo:     cof.TableNo(),
		ProductName: cof.ProductName(),
		CreatedAt:   cof.CreatedAt(),
	}

	err = ddbClient.Put(TableName, input)
	if err != nil {
		return fmt.Errorf("put item err in Save() %w", err)
	}

	return nil
}
