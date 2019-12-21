package order

import "time"

type EventItem struct {
	ProductID string
	Qty       int
	Price     int64
}

type Created struct {
	ID         string
	CreateTime time.Time
	TableNo    string
	Items      []EventItem
}

func NewCreatedEvent(id, tableNo string, items []Item, createTime time.Time) Created {

	eventItems := make([]EventItem, 0)
	for _, item := range items {
		eventItems = append(eventItems, EventItem{
			ProductID: item.productID,
			Qty:       item.qty,
			Price:     item.price,
		})
	}

	return Created{
		ID:         id,
		CreateTime: createTime,
		TableNo:    tableNo,
		Items:      eventItems,
	}
}

func (e Created) CreateAt() time.Time {
	return e.CreateTime
}

func (e Created) Identity() string {
	return e.ID
}
