package schematypes

const KeyEnumValue = "enumValue"

type EnumValue struct{}

var _ Type = (*EnumValue)(nil)

func (t *EnumValue) Key() string {
	return KeyEnumValue
}

func (t *EnumValue) String() string {
	return t.Key()
}
