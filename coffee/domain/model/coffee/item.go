package coffee

type Item struct {
	name string
}

func (i Item) Name() string {
	return i.name
}

func NewItem(name string) Item {
	return Item{name: name}
}
