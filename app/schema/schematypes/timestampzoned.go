package schematypes

const KeyTimestampZoned = "timestampZoned"

type TimestampZoned struct{}

var _ Type = (*TimestampZoned)(nil)

func (x *TimestampZoned) Key() string {
	return KeyTimestampZoned
}

func (x *TimestampZoned) String() string {
	return x.Key()
}

func (x *TimestampZoned) Sortable() bool {
	return true
}

func (x *TimestampZoned) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewTimestampZoned() *Wrapped {
	return Wrap(&TimestampZoned{})
}
