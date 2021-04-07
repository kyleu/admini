package schematypes

const KeyDate = "date"

type Date struct{}

var _ Type = (*Date)(nil)

func (t Date) Key() string {
	return KeyDate
}

func (t Date) String() string {
	return t.Key()
}
