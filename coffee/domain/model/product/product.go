package product

type ID string

type Product struct {
	id   ID
	name string
}

func (p Product) Name() string {
	return p.name
}

func (p Product) ID() ID {
	return p.id
}
