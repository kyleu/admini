package schematypes

const KeyInt = "int"

type Int struct {
	BitSize int
}

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

func NewInt(bitSize int) *Wrapped {
	return Wrap(&Int{BitSize: bitSize})
}
