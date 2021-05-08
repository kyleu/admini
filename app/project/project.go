package project

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
)

type Project struct {
	Key      string         `json:"key"`
	Title    string         `json:"title,omitempty"`
	Schemata []string       `json:"schemata"`
	Actions  action.Actions `json:"actions,omitempty"`
}

func (p Project) ModelsByPackage() *model.Package {
	return nil
}

type Projects []*Project

type View struct {
	Project  *Project                  `json:"project"`
	Schemata map[string]*schema.Schema `json:"schemata,omitempty"`
	Sources  map[string]*source.Source `json:"sources,omitempty"`
}
