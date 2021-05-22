package project

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
)

type Project struct {
	Key     string         `json:"-"`
	Title   string         `json:"title,omitempty"`
	Sources []string       `json:"sources,omitempty"`
	Actions action.Actions `json:"actions,omitempty"`
}

func (p *Project) Name() string {
	if p.Title == "" {
		return p.Key
	}
	return p.Title
}

func (p *Project) ModelsByPackage() *model.Package {
	return nil
}

type Projects []*Project

func (p Projects) Get(key string) *Project {
	for _, x := range p {
		if x.Key == key {
			return x
		}
	}
	return nil
}

type View struct {
	Project   *Project        `json:"project"`
	Schemata  schema.Schemata `json:"schemata,omitempty"`
	Sources   source.Sources  `json:"sources,omitempty"`
}
