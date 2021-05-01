package schematypes

const KeyInt = "int"

type Int struct{}

var _ Type = (*Int)(nil)

func (t *Int) Key() string {
	return KeyInt
}

func (t *Int) String() string {
	return t.Key()
}

func (t *Int) Sortable() bool {
	return true
}

func NewInt() *Wrapped {
	return Wrap(&Int{})
}
