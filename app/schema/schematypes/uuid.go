package schematypes

const KeyUUID = "uuid"

type UUID struct{}

var _ Type = (*UUID)(nil)

func (t *UUID) Key() string {
	return KeyUUID
}

func (t *UUID) String() string {
	return t.Key()
}
