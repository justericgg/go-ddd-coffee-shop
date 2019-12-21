package coffee

import (
	"github.com/justericgg/go-ddd-coffee-shop/ddd"
	"time"
)

type Coffee struct {
	ddd.AggregateRoot
	id        ID
	tableNo   string
	items     []Item
	createdAt time.Time
}

func Make(coffeeID ID, tableNo string, items []Item, createAt time.Time) *Coffee {

	//TODO: Verify by policy
	//TODO: Apply event

	return &Coffee{
		id:        coffeeID,
		tableNo:   tableNo,
		items:     items,
		createdAt: createAt,
	}
}
