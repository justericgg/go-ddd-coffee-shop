package repository

import "github.com/justericgg/go-ddd-coffee-shop/coffee/domain/model/product"

type ProductRepository struct{}

func (r *ProductRepository) GetProductList() (map[product.ID]string, error) {
	return map[product.ID]string{
		"A1": "Americano",
		"A2": "Latte",
	}, nil
}
