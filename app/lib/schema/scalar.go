// Content managed by Project Forge, see [projectforge.md] for details.
package schema

import (
	"admini.dev/admini/app/util"
)

type Scalar struct {
	Pkg         util.Pkg  `json:"pkg,omitempty"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	Description string    `json:"description,omitempty"`
	Metadata    *Metadata `json:"metadata,omitempty"`
}

type Scalars []*Scalar

func (s Scalars) Get(pkg util.Pkg, key string) *Scalar {
	for _, x := range s {
		if x.Pkg.Equals(pkg) && x.Key == key {
			return x
		}
	}
	return nil
}
