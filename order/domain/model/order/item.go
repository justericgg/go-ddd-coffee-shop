package order

type Item struct {
	productID string
	qty       int
	price     int64
}

func New(productID string, qty int, price int64) *Item {
	return &Item{
		productID: productID,
		qty:       qty,
		price:     price,
	}
}

func (i *Item) ProductID() string {
	return i.productID
}

func (i *Item) Qty() int {
	return i.qty
}

func (i *Item) Price() int64 {
	return i.price
}
