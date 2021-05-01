package schematypes

const KeyTimestamp = "timestamp"

type Timestamp struct{}

var _ Type = (*Timestamp)(nil)

func (t *Timestamp) Key() string {
	return KeyTimestamp
}

func (t *Timestamp) Sortable() bool {
	return true
}

func (t *Timestamp) String() string {
	return t.Key()
}

func NewTimestamp() *Wrapped {
	return Wrap(&Timestamp{})
}
