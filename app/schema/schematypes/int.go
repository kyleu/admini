package schematypes

import "fmt"

const KeyInt = "int"

type Int struct {
	Bits int `json:"bits,omitempty"`
}

var _ Type = (*Int)(nil)

func (t *Int) Key() string {
	return KeyInt
}

func (t *Int) String() string {
	if t.Bits > 0 {
		return fmt.Sprintf("%v%v", t.Key(), t.Bits)
	}
	return t.Key()
}

func (t *Int) Sortable() bool {
	return true
}

func NewInt(bits int) *Wrapped {
	return Wrap(&Int{Bits: bits})
}
