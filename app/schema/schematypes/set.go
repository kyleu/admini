package schematypes

import (
	"fmt"
)

const KeySet = "set"

type Set struct {
	T *Wrapped `json:"t"`
}

var _ Type = (*Set)(nil)

func (x *Set) Key() string {
	return KeySet
}

func (x *Set) String() string {
	return fmt.Sprintf("%s[%s]", x.Key(), x.T.String())
}

func (x *Set) Sortable() bool {
	return x.T.Sortable()
}

func (x *Set) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewSet() *Wrapped {
	return Wrap(&Set{})
}
