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

func (t *Reference) Key() string {
	return KeyReference
}

func (t *Reference) Sortable() bool {
	return true
}

func (t *Reference) String() string {
	return "ref:" + strings.Join(t.Path(), ".")
}

func (t *Reference) Path() []string {
	ret := make([]string, 0, len(t.Pkg)+1)
	copy(ret, t.Pkg)
	return append(ret, t.K)
}

func NewReference() *Wrapped {
	return Wrap(&Reference{})
}

func NewReferenceArgs(pkg util.Pkg, k string) *Wrapped {
	return Wrap(&Reference{Pkg: pkg, K: k})
}
