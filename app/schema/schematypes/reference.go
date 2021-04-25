package schematypes

import (
	"strings"
)

const KeyReference = "reference"

type Reference struct {
	Pkg []string `json:"pkg,omitempty"`
	T   string   `json:"t"`
}

func (t *Reference) Key() string {
	return KeyReference
}

func (t *Reference) String() string {
	return "ref:" + strings.Join(append(t.Pkg, t.T), ".")
}
