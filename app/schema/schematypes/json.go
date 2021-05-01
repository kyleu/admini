package schematypes

const KeyJSON = "json"

type JSON struct{}

var _ Type = (*JSON)(nil)

func (t *JSON) Key() string {
	return KeyJSON
}

func (t *JSON) String() string {
	return t.Key()
}

func (t *JSON) Sortable() bool {
	return false
}

func NewJSON() *Wrapped {
	return Wrap(&JSON{})
}
