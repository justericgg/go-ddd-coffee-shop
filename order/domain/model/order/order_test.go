package order_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	type args struct {
		id         order.ID
		tableNo    string
		status     order.Status
		items      []order.Item
		createDate time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "with empty item",
			args: args{
				id:         order.NewID(1, time.Now()),
				tableNo:    "",
				status:     0,
				items:      nil,
				createDate: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "apply order created event",
			args: args{
				id:         order.NewID(1, time.Now()),
				tableNo:    "123",
				status:     0,
				items:      []order.Item{order.NewItem("1", 1, 100)},
				createDate: time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := order.Create(tt.args.id, tt.args.tableNo, tt.args.status, tt.args.items, tt.args.createDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != nil {
				events := got.DomainEvents()
				_, ok := events[0].(order.Created)
				if !ok || len(events) == 0 {
					t.Errorf("Create() events got %v, want Created event", events)
				}
			}
		})
	}
}
