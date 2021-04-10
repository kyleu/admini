package schematypes

import (
	"fmt"
)

type Type interface {
	Key() string
	fmt.Stringer
}

type Wrapped struct {
	K string `json:"k"`
	T Type   `json:"t,omitempty"`
}

var _ Type = (*Wrapped)(nil)

func Wrap(t Type) Wrapped {
	_, ok := t.(Wrapped)
	if ok {
		return t.(Wrapped)
	}
	return Wrapped{K: t.Key(), T: t}
}

func (w Wrapped) Key() string {
	return w.K
}

func (w Wrapped) String() string {
	return w.T.String()
}
