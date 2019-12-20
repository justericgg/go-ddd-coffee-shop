package order

import (
	"github.com/justericgg/go-ddd-coffee-shop/ddd"
	"time"
)

type Order struct {
	ddd.AggregateRoot
	id         ID
	tableNo    string
	status     Status
	items      []Item
	createDate time.Time
	modifyDate time.Time
}

func Create(id ID, tableNo string, status Status, items []Item, createDate time.Time) (*Order, error) {
	order := &Order{
		id:         id,
		tableNo:    tableNo,
		status:     status,
		items:      items,
		createDate: createDate,
	}

	p := new(Policy)
	err := p.Verify(order)
	if err != nil {
		return nil, err
	}

	orderCreated := NewCreatedEvent(order.id.String(), order.tableNo, order.items)
	order.ApplyDomain(orderCreated)

	return order, nil
}
