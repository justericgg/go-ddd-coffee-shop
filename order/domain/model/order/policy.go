package order

import (
	"errors"
)

type Policy struct {
}

func (p *Policy) Verify(o *Order) error {
	idSpec := IdSpec{}
	if !idSpec.IsSatisfyBy(o.id) {
		return errors.New("order id spec invalid")
	}

	itemSpec := ItemSpec{}
	if !itemSpec.IsSatisfyBy(o.items) {
		return errors.New("at least one order item")
	}

	return nil
}
