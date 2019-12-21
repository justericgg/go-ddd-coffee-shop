package order

type Repository interface {
	GenerateID() ID
	Save(order Order)
}
