package coffee

import (
	"fmt"
	"time"
)

type ID struct {
	id string
}

func NewID(seq int64, createDate time.Time) ID {
	layout := "20060102"
	date := fmt.Sprintf(createDate.Format(layout))
	id := fmt.Sprintf("%s-%s-%d", "cof", date, seq)

	return ID{id: id}
}

func (id ID) String() string {
	return id.id
}
