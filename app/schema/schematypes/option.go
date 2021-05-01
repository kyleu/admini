package schematypes

const KeyOption = "option"

type Option struct {
	T *Wrapped `json:"t"`
}

var _ Type = (*Option)(nil)

func (t *Option) Key() string {
	return KeyOption
}

func (t *Option) Sortable() bool {
	return t.T.Sortable()
}

func (t *Option) String() string {
	return "*" + t.T.String()
}
