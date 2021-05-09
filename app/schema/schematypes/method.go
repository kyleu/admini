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

func (x *Method) Key() string {
	return KeyMethod
}

func (x *Method) String() string {
	argStrings := make([]string, 0, len(x.Args))
	for _, arg := range x.Args {
		argStrings = append(argStrings, arg.String())
	}
	return fmt.Sprintf("fn(%v) %v", strings.Join(argStrings, ", "), x.Ret.String())
}

func (x *Method) Sortable() bool {
	for _, a := range x.Args {
		if !a.Type.Sortable() {
			return false
		}
	}
	return x.Ret.Sortable()
}

func (x *Method) From(v interface{}) interface{} {
	switch t := v.(type) {
	default:
		return invalidInput(x.Key(), t)
	}
}

func NewMethod(ret *Wrapped) *Wrapped {
	return Wrap(&Method{Ret: ret})
}
