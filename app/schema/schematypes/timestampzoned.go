package schematypes

const KeyTimestampZoned = "timestampZoned"

type TimestampZoned struct{}

var _ Type = (*TimestampZoned)(nil)

func (t TimestampZoned) Key() string {
	return KeyTimestampZoned
}

func (t TimestampZoned) String() string {
	return t.Key()
}
