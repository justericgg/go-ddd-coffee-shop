// +build integration

package dynamo_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"testing"
	"time"
)

func TestClient_Count(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "empty count",
			args:    args{table: "Order"},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := dynamo.GetClient()
			if err != nil {
				t.Fatalf("Count() err %v", err)
			}
			got, err := c.Count(tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Put(t *testing.T) {
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
	type args struct {
		table string
		in    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "put item",
			args: args{
				table: "Order",
				in: Schema{
					OrderID: "ord-20191213-1",
					Status:  1,
					TableNo: "A12",
					Items: []Item{
						Item{
							ProductID: "product_id",
							Qty:       1,
							Price:     100,
						},
					},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := dynamo.GetClient()
			if err != nil {
				t.Fatalf("Put() err %v", err)
			}
			if err := c.Put(tt.args.table, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
