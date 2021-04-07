package schematypes

const KeyBool = "bool"

type Bool struct{}

var _ Type = (*Bool)(nil)

func (t Bool) Key() string {
	return KeyBool
}

func (t Bool) String() string {
	return t.Key()
}
