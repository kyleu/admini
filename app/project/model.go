package project

import (
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
)

type Project struct {
	Key string
	Title string
	Schemata []string
}

func (p Project) ModelsByPackage() *schema.ModelPackage {
	return nil
}

type Projects []*Project

type View struct {
	Project *Project
	Schemata map[string]*schema.Schema
	Sources map[string]*source.Source
}
