package schematypes

const KeyXML = "xml"

type XML struct{}

var _ Type = (*XML)(nil)

func (t XML) Key() string {
	return KeyXML
}

func (t XML) String() string {
	return t.Key()
}
