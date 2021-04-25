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
