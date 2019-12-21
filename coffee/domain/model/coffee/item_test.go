package coffee

import (
	"testing"
)

func TestNewItem(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name         string
		args         args
		wantItemName string
	}{
		{
			name:         "set right item name",
			args:         args{name: "Americano"},
			wantItemName: "Americano",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewItem(tt.args.name)
			if got.Name() != tt.wantItemName {
				t.Errorf("NewItem().name = %s, want %s", got.Name(), tt.wantItemName)
			}
		})
	}
}
