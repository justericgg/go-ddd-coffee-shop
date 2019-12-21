package order

type Repository interface {
	GenerateID() (ID, error)
	Save(order *Order) error
}
