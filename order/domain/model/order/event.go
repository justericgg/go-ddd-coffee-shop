package order

import "time"

type Created struct {
	ID         string
	CreateTime time.Time
	TableNo    string
	Items      []Item
}

func NewCreatedEvent(id, tableNo string, items []Item) Created {
	return Created{
		ID:         id,
		CreateTime: time.Now(),
		TableNo:    tableNo,
		Items:      items,
	}
}

func (e Created) CreateAt() time.Time {
	return e.CreateTime
}

func (e Created) Identity() string {
	return e.ID
}
