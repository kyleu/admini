package schematypes

const KeyChar = "char"

type Char struct{}

var _ Type = (*Char)(nil)

func (x *Char) Key() string {
	return KeyChar
}

func (x *Char) String() string {
	return x.Key()
}

func (x *Char) Sortable() bool {
	return true
}

func (x *Char) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewChar() *Wrapped {
	return Wrap(&Char{})
}
