package schematypes

const KeyByte = "byte"

type Byte struct{}

var _ Type = (*Byte)(nil)

func (x *Byte) Key() string {
	return KeyByte
}

func (x *Byte) String() string {
	return x.Key()
}

func (x *Byte) Sortable() bool {
	return true
}

func (x *Byte) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewByte() *Wrapped {
	return Wrap(&Byte{})
}
