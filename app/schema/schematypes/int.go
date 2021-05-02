package schematypes

import "fmt"

const KeyInt = "int"

type Int struct {
	BitSize int
}

var _ Type = (*Int)(nil)

func (t *Int) Key() string {
	return KeyInt
}

func (t *Int) String() string {
	if t.BitSize > 0 {
		return fmt.Sprintf("%v%v", t.Key(), t.BitSize)
	}
	return t.Key()
}

func (t *Int) Sortable() bool {
	return true
}

func NewInt(bitSize int) *Wrapped {
	return Wrap(&Int{BitSize: bitSize})
}
