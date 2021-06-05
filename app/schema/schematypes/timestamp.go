package schematypes

const KeyTimestamp = "timestamp"

type Timestamp struct{}

var _ Type = (*Timestamp)(nil)

func (x *Timestamp) Key() string {
	return KeyTimestamp
}

func (x *Timestamp) String() string {
	return x.Key()
}

func (x *Timestamp) Sortable() bool {
	return true
}

func (x *Timestamp) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewTimestamp() *Wrapped {
	return Wrap(&Timestamp{})
}
