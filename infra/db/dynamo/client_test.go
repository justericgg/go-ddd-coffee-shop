// +build integration

package dynamo_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/infra/db/dynamo"
	"testing"
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
