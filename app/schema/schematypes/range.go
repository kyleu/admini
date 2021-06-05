package schematypes

import (
	"fmt"
)

const KeyRange = "range"

var _ Type = (*Range)(nil)

type Range struct {
	T *Wrapped `json:"t"`
}

func (x *Range) Key() string {
	return KeyRange
}

func (x *Range) String() string {
	return fmt.Sprintf("range[%s]", x.T.String())
}

func (x *Range) Sortable() bool {
	return x.T.Sortable()
}

func (x *Range) From(v interface{}) interface{} {
	return invalidInput(x.Key(), v)
}

func NewRange(t *Wrapped) *Wrapped {
	return Wrap(&Range{T: t})
}
