package schematypes

const KeyUUID = "uuid"

type UUID struct{}

var _ Type = (*UUID)(nil)

func (x *UUID) Key() string {
	return KeyUUID
}

func (x *UUID) Sortable() bool {
	return true
}

func (x *UUID) String() string {
	return x.Key()
}

func (x *UUID) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewUUID() *Wrapped {
	return Wrap(&UUID{})
}
