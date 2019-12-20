package order_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"testing"
)

func TestNewItem(t *testing.T) {
	type args struct {
		productID string
		qty       int
		price     int64
	}
	tests := []struct {
		name          string
		args          args
		wantProductID string
		wantQty       int
		wantPrice     int64
	}{
		{
			name: "getter test",
			args: args{
				productID: "productID",
				qty:       1,
				price:     100,
			},
			wantProductID: "productID",
			wantQty:       1,
			wantPrice:     100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := order.NewItem(tt.args.productID, tt.args.qty, tt.args.price)
			if got.ProductID() != tt.wantProductID {
				t.Errorf("NewItem().ProductID = %s, want %s", got.ProductID(), tt.wantProductID)
			}
			if got.Qty() != tt.wantQty {
				t.Errorf("NewItem().ProductID = %d, want %d", got.Qty(), tt.wantQty)
			}
			if got.Price() != tt.wantPrice {
				t.Errorf("NewItem().ProductID = %d, want %d", got.Price(), tt.wantPrice)
			}
		})
	}
}
