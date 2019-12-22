package coffee

import (
	"github.com/justericgg/go-ddd-coffee-shop/ddd"
	"time"
)

type Coffee struct {
	ddd.AggregateRoot
	id          ID
	orderID     string
	tableNo     string
	productName string
	createdAt   time.Time
}

func (c *Coffee) OrderID() string {
	return c.orderID
}

func (c *Coffee) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Coffee) ProductName() string {
	return c.productName
}

func (c *Coffee) TableNo() string {
	return c.tableNo
}

func (c *Coffee) ID() ID {
	return c.id
}

func Make(coffeeID ID, orderID, tableNo, productName string, createAt time.Time) *Coffee {

	//TODO: Verify by policy
	//TODO: Apply event

	return &Coffee{
		id:          coffeeID,
		orderID:     orderID,
		tableNo:     tableNo,
		productName: productName,
		createdAt:   createAt,
	}
}
