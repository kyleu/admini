package schematypes

import "fmt"

const KeyInt = "int"

type Int struct {
	Bits int `json:"bits,omitempty"`
}

var _ Type = (*Int)(nil)

func (x *Int) Key() string {
	return KeyInt
}

func (x *Int) String() string {
	if x.Bits > 0 {
		return fmt.Sprintf("%v%v", x.Key(), x.Bits)
	}
	return x.Key()
}

func (x *Int) Sortable() bool {
	return true
}

func (x *Int) From(v interface{}) interface{} {
	switch t := v.(type) {
	case int:
		return t
	case int32:
		return t
	case int64:
		return t
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewInt(bits int) *Wrapped {
	return Wrap(&Int{Bits: bits})
}
