package schematypes

import (
	"strings"

	"github.com/kyleu/admini/app/util"
)

const KeyReference = "reference"

type Reference struct {
	Pkg util.Pkg `json:"pkg,omitempty"`
	T   string   `json:"t"`
}

func (t *Reference) Key() string {
	return KeyReference
}

func (t *Reference) String() string {
	return strings.Join(append(t.Pkg, t.T), ".")
}
