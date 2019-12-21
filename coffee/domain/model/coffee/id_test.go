package coffee_test

import (
	"github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/coffee"
	"testing"
	"time"
)

func TestNewID(t *testing.T) {
	type args struct {
		seq        int64
		createDate time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "must be the right format",
			args: args{
				seq:        1,
				createDate: time.Date(2019, 1, 1, 12, 13, 1, 0, time.UTC),
			},
			want: "cof-20190101-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coffee.NewID(tt.args.seq, tt.args.createDate)
			if got.String() != tt.want {
				t.Errorf("NewID() = %v, want %v", got, tt.want)
			}
		})
	}
}
