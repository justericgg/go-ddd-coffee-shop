package order_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"testing"
)

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    order.Status
		want string
	}{
		{
			name: "Initial",
			s:    order.Initial,
			want: "Initial",
		},
		{
			name: "Processing",
			s:    order.Processing,
			want: "Processing",
		},
		{
			name: "Deliver",
			s:    order.Deliver,
			want: "Deliver",
		},
		{
			name: "Closed",
			s:    order.Closed,
			want: "Closed",
		},
		{
			name: "Cancel",
			s:    order.Cancel,
			want: "Cancel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
