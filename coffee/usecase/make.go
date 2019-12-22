package usecase

import (
	"fmt"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/coffee"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/product"
	"github.com/justericgg/go-ddd-coffee-shop/coffee/infra/repository"
	"time"
)

type Item struct {
	ProductID string
	Qty       int
}

type MakeCoffeeCmd struct {
	TableNo string
	OrderID string
	Items   []Item
}

type MakeCoffeeResult struct {
	CoffeeID    string
	ProductID   string
	ProductName string
}

type MakeCoffeeSvc struct {
	coffeeRepo  repository.CoffeeRepository
	productRepo repository.ProductRepository
}

func NewMakeCoffeeSvc(c repository.CoffeeRepository, p repository.ProductRepository) *MakeCoffeeSvc {
	return &MakeCoffeeSvc{
		coffeeRepo:  c,
		productRepo: p,
	}
}

func (s *MakeCoffeeSvc) MakeCoffee(cmd MakeCoffeeCmd) ([]MakeCoffeeResult, error) {
	productList, err := s.productRepo.GetProductList()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	results := make([]MakeCoffeeResult, 0, len(cmd.Items))
	for _, item := range cmd.Items {
		productName, ok := productList[product.ID(item.ProductID)]
		if !ok {
			return nil, fmt.Errorf("product id %s not exists in list %v", item.ProductID, productList)
		}
		coffeeID, err := s.coffeeRepo.GenerateID()
		if err != nil {
			return nil, fmt.Errorf("generate coffee id err in MakeCoffee %w", err)
		}
		c := coffee.Make(coffeeID, cmd.OrderID, cmd.TableNo, productName, now)
		err = s.coffeeRepo.Save(c)
		if err != nil {
			return nil, fmt.Errorf("save coffee err in MakeCoffee, %w", err)
		}
		results = append(results, MakeCoffeeResult{
			CoffeeID:    c.ID().String(),
			ProductID:   item.ProductID,
			ProductName: c.ProductName(),
		})
	}

	return results, nil
}
