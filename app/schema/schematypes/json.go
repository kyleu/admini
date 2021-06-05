package schematypes

const KeyJSON = "json"

type JSON struct{}

var _ Type = (*JSON)(nil)

func (x *JSON) Key() string {
	return KeyJSON
}

func (x *JSON) String() string {
	return x.Key()
}

func (x *JSON) Sortable() bool {
	return false
}

func (x *JSON) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewJSON() *Wrapped {
	return Wrap(&JSON{})
}
