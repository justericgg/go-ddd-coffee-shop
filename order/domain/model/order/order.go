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

func (o *Order) ModifyDate() time.Time {
	return o.modifyDate
}

func (o *Order) CreateDate() time.Time {
	return o.createDate
}

func (o *Order) Items() []Item {
	return o.items
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) TableNo() string {
	return o.tableNo
}

func (o *Order) ID() ID {
	return o.id
}

func Create(id ID, tableNo string, items []Item, createDate time.Time) (*Order, error) {
	order := &Order{
		id:         id,
		tableNo:    tableNo,
		status:     Initial,
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
