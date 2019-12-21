package usecase

import (
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

type CreateOrderSvc struct {
	repository     order.Repository
	eventPublisher EventPublisher
}

type CreateOrderCmd struct {
	TableNo string
	Items   []struct {
		ProductID string
		Qty       int
		Price     int64
	}
}

func (s *CreateOrderSvc) CreateOrder(cmd CreateOrderCmd) error {
	orderID := s.repository.GenerateID()
	orderItem := make([]order.Item, 0)
	for _, item := range cmd.Items {
		orderItem = append(orderItem, order.NewItem(item.ProductID, item.Qty, item.Price))
	}

	ord, err := order.Create(orderID, cmd.TableNo, orderItem, time.Now())
	if err != nil {
		return err
	}

	s.eventPublisher.Publish(ord.DomainEvents())

	return nil
}
