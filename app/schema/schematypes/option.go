package schematypes

const KeyOption = "option"

type Option struct {
	T *Wrapped `json:"t"`
}

var _ Type = (*Option)(nil)

func (x *Option) Key() string {
	return KeyOption
}

func (x *Option) String() string {
	return "*" + x.T.String()
}

func (x *Option) Sortable() bool {
	return x.T.Sortable()
}

func (x *Option) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewOption(t *Wrapped) *Wrapped {
	return Wrap(&Option{T: t})
}
