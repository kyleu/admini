package schematypes

import (
	"strings"

	"github.com/kyleu/admini/app/util"
)

const KeyReference = "reference"

type Reference struct {
	Pkg []string `json:"pkg,omitempty"`
	K   string   `json:"k"`
}

func (x *Reference) Key() string {
	return KeyReference
}

func (x *Reference) Sortable() bool {
	return true
}

func (x *Reference) String() string {
	return "ref:" + strings.Join(x.Path(), ".")
}

func (x *Reference) From(v interface{}) interface{} {
	return invalidInput(x.Key(), x)
}

func (x *Reference) Path() []string {
	ret := util.StringArrayCopy(x.Pkg)
	return append(ret, x.K)
}

func NewReference() *Wrapped {
	return Wrap(&Reference{})
}

func NewReferenceArgs(pkg util.Pkg, k string) *Wrapped {
	return Wrap(&Reference{Pkg: pkg, K: k})
}
