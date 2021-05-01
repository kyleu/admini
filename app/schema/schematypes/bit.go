package schematypes

const KeyBit = "bit"

type Bit struct{}

var _ Type = (*Bit)(nil)

func (t *Bit) Key() string {
	return KeyBit
}

func (t *Bit) String() string {
	return t.Key()
}

func (t *Bit) Sortable() bool {
	return true
}

func NewBit() *Wrapped {
	return Wrap(&Bit{})
}
