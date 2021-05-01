package schematypes

const KeyTimestampZoned = "timestampZoned"

type TimestampZoned struct{}

var _ Type = (*TimestampZoned)(nil)

func (t *TimestampZoned) Key() string {
	return KeyTimestampZoned
}

func (t *TimestampZoned) Sortable() bool {
	return true
}

func (t *TimestampZoned) String() string {
	return t.Key()
}

func NewTimestampZoned() *Wrapped {
	return Wrap(&TimestampZoned{})
}
