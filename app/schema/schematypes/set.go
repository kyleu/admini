package schematypes

import (
	"fmt"
)

const KeySet = "set"

type Set struct {
	T Wrapped
}

var _ Type = (*Set)(nil)

func (t Set) Key() string {
	return KeySet
}

func (t Set) String() string {
	return fmt.Sprintf("%v[%v]", t.Key(), t.T.String())
}
