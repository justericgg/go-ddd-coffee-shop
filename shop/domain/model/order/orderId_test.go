package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeOrderIdAndGetOrderIdValue(t *testing.T) {

	orderId := MakeOrderId("order_id")

	assert.Equal(t, "order_id", orderId.toString(), "must get a order id string value")
}
