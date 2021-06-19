package schematypes

const KeyOption = "option"

type Option struct {
	V *Wrapped `json:"v"`
}

var _ Type = (*Option)(nil)

func (x *Option) Key() string {
	return KeyOption
}

func (x *Option) String() string {
	return "*" + x.V.String()
}

func (x *Option) Sortable() bool {
	return x.V.Sortable()
}

func (x *Option) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func NewOption(t *Wrapped) *Wrapped {
	return Wrap(&Option{V: t})
}

func (x *Option) Default(string) interface{} {
	return "∅"
}
