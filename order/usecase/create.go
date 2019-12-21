package usecase

import (
	"github.com/justericgg/go-ddd-coffee-shop/order/domain/model/order"
	"time"
)

type CreateOrderSvc struct {
	repository     order.Repository
	eventPublisher EventPublisher
}

func NewCreateOrderSvc(r order.Repository, publisher EventPublisher) *CreateOrderSvc {
	return &CreateOrderSvc{
		repository:     r,
		eventPublisher: publisher,
	}
}

type Item struct {
	ProductID string
	Qty       int
	Price     int64
}
type CreateOrderCmd struct {
	TableNo string
	Items   []Item
}

func (s *CreateOrderSvc) CreateOrder(cmd CreateOrderCmd) error {
	orderID, err := s.repository.GenerateID()
	if err != nil {
		return err
	}
	orderItem := make([]order.Item, 0)
	for _, item := range cmd.Items {
		orderItem = append(orderItem, order.NewItem(item.ProductID, item.Qty, item.Price))
	}

	ord, err := order.Create(orderID, cmd.TableNo, orderItem, time.Now())
	if err != nil {
		return err
	}

	err = s.repository.Save(ord)
	if err != nil {
		return err
	}

	s.eventPublisher.Publish(ord.DomainEvents())

	return nil
}
