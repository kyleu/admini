package schematypes

import (
	"fmt"
)

const KeyList = "list"

var _ Type = (*List)(nil)

type List struct {
	T *Wrapped `json:"t"`
}

func (t *List) Key() string {
	return KeyList
}

func (t *List) String() string {
	return fmt.Sprintf("[]%v", t.T.String())
}
