package project

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
)

const SourceProjectPrefix = "__"

func (s *Service) LoadSourceProject(srcKey string) (*View, error) {
	src, err := s.sources.Load(srcKey, false)
	if err != nil {
		return nil, err
	}
	sch, err := s.sources.LoadSchema(src.Key)
	if err != nil {
		return nil, err
	}
	if sch == nil {
		return nil, errors.Errorf("no schema found with key [%s]", srcKey)
	}

	sources := source.Sources{src}
	schemata := schema.Schemata{src.Key: sch}

	acts := make(action.Actions, 0, len(sch.Models))

	mp := sch.ModelsByPackage()

	if len(mp.ChildPackages) > 0 {
		for _, p := range mp.ChildPackages {
			pCfg := util.ValueMapFor(action.TypeSource.Key, src.Key, action.TypePackage.Key, p.Key)
			acts = append(acts, &action.Action{Key: p.Key, Type: action.TypePackage, Title: p.Name(), Config: pCfg})
		}
		acts = append(acts, &action.Action{Key: action.TypeSeparator.Key, Type: action.TypeSeparator})
	}

	if len(mp.ChildModels) > 0 {
		for _, m := range mp.ChildModels {
			mCfg := util.ValueMapFor(action.TypeSource.Key, src.Key, action.TypeModel.Key, strings.Join(append(m.Pkg, m.Key), "/"))
			acts = append(acts, &action.Action{Key: m.Key, Type: action.TypeModel, Title: m.Name(), Config: mCfg})
		}
		acts = append(acts, &action.Action{Key: action.TypeSeparator.Key, Type: action.TypeSeparator})
	}

	sCfg := util.ValueMapFor(action.TypeSource.Key, src.Key, action.TypeActivity.Key, "sql")
	sd := "Run ad-hoc SQL queries and DDL"
	acts = append(acts, &action.Action{Key: "sql", Type: action.TypeActivity, Title: "SQL Playground", Description: sd, Ordinal: 0, Config: sCfg})

	prj := &Project{Key: SourceProjectPrefix + src.Key, Title: src.Title, Icon: src.Icon, Description: src.Description, Sources: []string{srcKey}, Actions: acts}

	return &View{Project: prj, Sources: sources, Schemata: schemata}, nil
}
