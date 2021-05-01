package schematypes

const KeyXML = "xml"

type XML struct{}

var _ Type = (*XML)(nil)

func (t *XML) Key() string {
	return KeyXML
}

func (t *XML) Sortable() bool {
	return true
}

func (t *XML) String() string {
	return t.Key()
}

func NewXML() *Wrapped {
	return Wrap(&XML{})
}
