package repository

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/coffee"
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

const tableName = "Coffee"

type Schema struct {
	CoffeeID    string
	TableNo     string
	ProductName string
	CreatedAt   time.Time
}

type CoffeeRepository struct{}

func (c *CoffeeRepository) GenerateID() (coffee.ID, error) {
	ddbClient, err := dynamo.GetClient()
	if err != nil {
		return coffee.ID{}, fmt.Errorf("get conn err in GenerateID() %w", err)
	}
	count, err := ddbClient.Count(tableName)
	if err != nil {
		return coffee.ID{}, fmt.Errorf("count err in GenerateID() %w", err)
	}

	id := order.NewID(count+1, time.Now())

	return id, nil
}

func (c *CoffeeRepository) Save(cof *coffee.Coffee) error {
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

	err = ddbClient.Put(tableName, input)
	if err != nil {
		return fmt.Errorf("put item err in Save() %w", err)
	}

	return nil
}
