// Content managed by Project Forge, see [projectforge.md] for details.
package model

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"

	"admini.dev/admini/app/util"
)

type Models []*Model

func (m Models) Get(pkg util.Pkg, key string) *Model {
	return lo.FindOrElse(m, nil, func(x *Model) bool {
		return x.Pkg.Equals(pkg) && x.Key == key
	})
}

func (m Models) Sort() {
	slices.SortFunc(m, func(l *Model, r *Model) bool {
		return l.Key < r.Key
	})
}

func (m Models) Names() []string {
	return lo.Map(m, func(x *Model, _ int) string {
		return x.Key
	})
}
