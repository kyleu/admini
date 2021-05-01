package schematypes

const KeyBool = "bool"

type Bool struct{}

var _ Type = (*Bool)(nil)

func (t *Bool) Key() string {
	return KeyBool
}

func (t *Bool) String() string {
	return t.Key()
}

func (t *Bool) Sortable() bool {
	return true
}

func NewBool() *Wrapped {
	return Wrap(&Bool{})
}
