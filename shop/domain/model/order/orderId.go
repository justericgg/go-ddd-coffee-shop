package order

type Id struct {
	id string
}

func MakeOrderId(id string) Id {
	return Id{id: id}
}

func (id *Id) toString() string {
	return id.id
}
