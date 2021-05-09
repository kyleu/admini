package schematypes

const KeyXML = "xml"

type XML struct{}

var _ Type = (*XML)(nil)

func (x *XML) Key() string {
	return KeyXML
}

func (x *XML) String() string {
	return x.Key()
}

func (x *XML) Sortable() bool {
	return true
}

func (x *XML) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewXML() *Wrapped {
	return Wrap(&XML{})
}
