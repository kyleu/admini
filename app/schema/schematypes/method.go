package schematypes

import (
	"fmt"
	"strings"
)

type Argument struct {
	Key  string   `json:"key"`
	Type *Wrapped `json:"type"`
}

func (a Argument) String() string {
	return fmt.Sprintf("%v %v", a.Key, a.Type.String())
}

type Arguments []Argument

const KeyMethod = "method"

type Method struct {
	Args Arguments `json:"args,omitempty"`
	Ret  *Wrapped  `json:"ret,omitempty"`
}

var _ Type = (*Method)(nil)

func (t *Method) Key() string {
	return KeyMethod
}

func (t *Method) String() string {
	argStrings := make([]string, 0, len(t.Args))
	for _, arg := range t.Args {
		argStrings = append(argStrings, arg.String())
	}
	return fmt.Sprintf("fn(%v) %v", strings.Join(argStrings, ", "), t.Ret.String())
}

func (t *Method) Sortable() bool {
	for _, a := range t.Args {
		if !a.Type.Sortable() {
			return false
		}
	}
	return t.Ret.Sortable()
}

func NewMethod(ret *Wrapped) *Wrapped {
	return Wrap(&Method{Ret: ret})
}
