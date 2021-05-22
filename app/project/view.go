package project

import (
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
)

type View struct {
	Project   *Project        `json:"project"`
	Schemata  schema.Schemata `json:"schemata,omitempty"`
	Sources   source.Sources  `json:"sources,omitempty"`
}
