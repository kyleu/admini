package schematypes

const KeyEnumValue = "enumValue"

type EnumValue struct{}

var _ Type = (*EnumValue)(nil)

func (x *EnumValue) Key() string {
	return KeyEnumValue
}

func (x *EnumValue) String() string {
	return x.Key()
}

func (x *EnumValue) Sortable() bool {
	return true
}

func (x *EnumValue) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewEnumValue() *Wrapped {
	return Wrap(&EnumValue{})
}
