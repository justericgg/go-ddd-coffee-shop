package usecase

import "github.com/justericgg/go-ddd-coffee-shop/ddd"

type EventPublisher interface {
	Publish([]ddd.DomainEvent) error
}
