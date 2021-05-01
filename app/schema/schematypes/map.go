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

func (t *Map) Key() string {
	return KeyMap
}

func (t *Map) String() string {
	return fmt.Sprintf("%v[%v]%v", t.Key(), t.K.String(), t.V.String())
}

func (t *Map) Sortable() bool {
	return t.K.Sortable() && t.V.Sortable()
}

func NewMap(k *Wrapped, v *Wrapped) *Wrapped {
	return Wrap(&Map{K: k, V: v})
}
