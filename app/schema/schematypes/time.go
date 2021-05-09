package schematypes

const KeyTime = "time"

type Time struct{}

var _ Type = (*Time)(nil)

func (x *Time) Key() string {
	return KeyTime
}

func (x *Time) Sortable() bool {
	return true
}

func (x *Time) String() string {
	return x.Key()
}

func (x *Time) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewTime() *Wrapped {
	return Wrap(&Time{})
}
