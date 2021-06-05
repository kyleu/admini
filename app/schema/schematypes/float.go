package schematypes

const KeyFloat = "float"

type Float struct {
	Bits int
}

var _ Type = (*Float)(nil)

func (x *Float) Key() string {
	return KeyFloat
}

func (x *Float) String() string {
	return x.Key()
}

func (x *Float) Sortable() bool {
	return true
}

func (x *Float) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewFloat(bits int) *Wrapped {
	return Wrap(&Float{Bits: bits})
}
