package schematypes

const KeyBit = "bit"

type Bit struct{}

var _ Type = (*Bit)(nil)

func (x *Bit) Key() string {
	return KeyBit
}

func (x *Bit) String() string {
	return x.Key()
}

func (x *Bit) Sortable() bool {
	return true
}

func (x *Bit) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewBit() *Wrapped {
	return Wrap(&Bit{})
}
