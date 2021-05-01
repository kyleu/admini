package schematypes

const KeyOption = "option"

type Option struct {
	T *Wrapped `json:"t"`
}

var _ Type = (*Option)(nil)

func (o *Option) Key() string {
	return KeyOption
}

func (t *Option) Sortable() bool {
	return t.T.Sortable()
}

func (o *Option) String() string {
	return "*" + o.T.String()
}
