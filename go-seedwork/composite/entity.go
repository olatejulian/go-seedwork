package composite

import vo "github.com/olatejulian/go-seedwork/value_object"

type Entity struct {
	ID        vo.ID
	CreatedAt string
	UpdatedAt int
}

func (e *Entity) ChangeUpdateDate() {
	e.UpdatedAt = 2
}
