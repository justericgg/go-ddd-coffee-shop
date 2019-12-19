package order

import "testing"

func TestStatus_String(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want string
	}{
		{
			name: "Initial",
			s:    Initial,
			want: "Initial",
		},
		{
			name: "Processing",
			s:    Processing,
			want: "Processing",
		},
		{
			name: "Deliver",
			s:    Deliver,
			want: "Deliver",
		},
		{
			name: "Closed",
			s:    Closed,
			want: "Closed",
		},
		{
			name: "Cancel",
			s:    Cancel,
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
