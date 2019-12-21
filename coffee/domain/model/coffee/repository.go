package coffee

type Repository interface {
	GenerateID() (ID, error)
	Save(coffee Coffee) error
}
