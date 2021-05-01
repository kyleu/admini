package schematypes

const KeyFloat = "float"

type Float struct{}

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

func NewFloat() *Wrapped {
	return Wrap(&Float{})
}
