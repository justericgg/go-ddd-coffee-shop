package product

type Repository interface {
	GetProductList() (map[ID]string, error)
}
