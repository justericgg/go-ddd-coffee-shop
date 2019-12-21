package repository

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

const orderTableName = "Order"

type Item struct {
	ProductID string
	Qty       int
	Price     int64
}

type Schema struct {
	OrderID   string
	Status    int
	TableNo   string
	Items     []Item
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

func (o OrderRepository) Save(ord order.Order) error {
	ddbClient, err := dynamo.GetClient()
	if err != nil {
		return fmt.Errorf("get conn err in Save() %w", err)
	}

	input := &Schema{
		OrderID:   ord.ID().String(),
		Status:    int(ord.Status()),
		TableNo:   ord.TableNo(),
		CreatedAt: ord.CreateDate(),
		UpdatedAt: ord.ModifyDate(),
	}

	for _, item := range ord.Items() {
		input.Items = append(input.Items, Item{
			ProductID: item.ProductID(),
			Qty:       item.Qty(),
			Price:     item.Price(),
		})
	}

	err = ddbClient.Put(orderTableName, input)
	if err != nil {
		return fmt.Errorf("put item err in Save() %w", err)
	}

	return nil
}
