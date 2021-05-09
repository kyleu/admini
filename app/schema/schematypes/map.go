package schematypes

import (
	"fmt"
)

const KeyMap = "map"

var _ Type = (*Map)(nil)

type Map struct {
	K *Wrapped `json:"k"`
	V *Wrapped `json:"v"`
}

func (x *Map) Key() string {
	return KeyMap
}

func (x *Map) String() string {
	return fmt.Sprintf("%v[%v]%v", x.Key(), x.K.String(), x.V.String())
}

func (x *Map) Sortable() bool {
	return x.K.Sortable() && x.V.Sortable()
}

func (x *Map) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewMap(k *Wrapped, v *Wrapped) *Wrapped {
	return Wrap(&Map{K: k, V: v})
}
