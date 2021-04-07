package schematypes

const KeyTime = "time"

type Time struct{}

var _ Type = (*Time)(nil)

func (t Time) Key() string {
	return KeyTime
}

func (t Time) String() string {
	return t.Key()
}
