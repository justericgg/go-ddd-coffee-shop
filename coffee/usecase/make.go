package usecase

import (
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
	tableNo string
	Items   []Item
}

type MakeCoffeeSvc struct {
	coffeeRepo  repository.CoffeeRepository
	productRepo repository.ProductRepository
}

func (s *MakeCoffeeSvc) MakeCoffee(cmd MakeCoffeeCmd) error {
	productList, err := s.productRepo.GetProductList()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, item := range cmd.Items {
		productName := productList[product.ID(item.ProductID)]
		coffeeID, err := s.coffeeRepo.GenerateID()
		if err != nil {
			return err
		}
		c := coffee.Make(coffeeID, cmd.tableNo, productName, now)
		err = s.coffeeRepo.Save(c)
		if err != nil {
			return err
		}
	}

	return nil
}
