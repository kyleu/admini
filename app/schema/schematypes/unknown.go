package schematypes

const KeyUnknown = "unknown"

type Unknown struct {
	X string `json:"x"`
}

var _ Type = (*Unknown)(nil)

func (t *Unknown) Key() string {
	return KeyUnknown
}

func (t *Unknown) Sortable() bool {
	return false
}

func (t *Unknown) String() string {
	return t.Key() + "(" + t.X + ")"
}

func NewUnknown(x string) *Wrapped {
	return Wrap(&Unknown{X: x})
}
