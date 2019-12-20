package order

import (
	"regexp"
)

type IdSpec struct {
}

func (s IdSpec) IsSatisfyBy(orderID ID) bool {
	match, _ := regexp.MatchString(`ord-\d{8}-\d{1,}`, orderID.String())
	if !match {
		return false
	}
	return true
}

type ItemSpec struct {
}

func (s ItemSpec) IsSatisfyBy(item []Item) bool {
	if item == nil {
		return false
	}

	return true
}

type TableNoSpec struct {
}

type StatusTransitionSpec struct {
}
