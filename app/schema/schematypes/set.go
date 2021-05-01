package schematypes

import (
	"fmt"
)

const KeySet = "set"

type Set struct {
	T *Wrapped `json:"t"`
}

var _ Type = (*Set)(nil)

func (t *Set) Key() string {
	return KeySet
}

func (t *Set) Sortable() bool {
	return t.T.Sortable()
}

func (t *Set) String() string {
	return fmt.Sprintf("%v[%v]", t.Key(), t.T.String())
}
