package event

import (
	"github.com/justericgg/go-ddd-coffee-shop/ddd"
	"log"
)

type Cwe struct{}

func (cwe Cwe) Publish(events []ddd.DomainEvent) {
	for _, event := range events {
		log.Printf("event published: %v\n", event)
	}
}
