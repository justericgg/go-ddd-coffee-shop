package order

type Status int

const (
	Initial Status = 1 + iota
	Processing
	Deliver
	Closed
	Cancel
)

func (s Status) String() string {
	return [...]string{"Initial", "Processing", "Deliver", "Closed", "Cancel"}[s-1]
}
