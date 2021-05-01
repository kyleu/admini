package schematypes

import (
	"fmt"
)

const KeyRange = "range"

var _ Type = (*Range)(nil)

type Range struct {
	T *Wrapped `json:"t"`
}

func (t *Range) Key() string {
	return KeyRange
}

func (t *Range) String() string {
	return fmt.Sprintf("range[%v]", t.T.String())
}

func (t *Range) Sortable() bool {
	return t.T.Sortable()
}

func NewRange(t *Wrapped) *Wrapped {
	return Wrap(&Range{T: t})
}
