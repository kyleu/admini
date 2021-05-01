package schematypes

const KeyFloat = "float"

type Float struct {
	BitSize int
}

var _ Type = (*Float)(nil)

func (t *Float) Key() string {
	return KeyFloat
}

func (t *Float) String() string {
	return t.Key()
}

func (t *Float) Sortable() bool {
	return true
}

func NewFloat(bitSize int) *Wrapped {
	return Wrap(&Float{BitSize: bitSize})
}
