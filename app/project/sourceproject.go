package project

import (
	"strings"

	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"
)

const SourceProjectPrefix = "__"

func (s *Service) LoadSourceProject(sourceKey string) (*View, error) {
	src, err := s.sources.Load(sourceKey, false)
	if err != nil {
		return nil, err
	}
	sch, err := s.sources.LoadSchema(src.Key)
	if err != nil {
		return nil, err
	}

	sources := source.Sources{src}
	schemata := schema.Schemata{src.Key: sch}

	acts := make(action.Actions, 0, len(sch.Models))

	mp := sch.ModelsByPackage()

	if len(mp.ChildPackages) > 0 {
		for _, p := range mp.ChildPackages {
			pCfg := map[string]string{"source": src.Key, "package": p.Key}
			acts = append(acts, &action.Action{Key: p.Key, Type: "package", Title: p.Key, Config: pCfg})
		}
		acts = append(acts, &action.Action{Key: "separator", Type: "separator"})
	}

	if len(mp.ChildModels) > 0 {
		for _, m := range mp.ChildModels {
			mCfg := map[string]string{"source": src.Key, "model": strings.Join(append(m.Pkg, m.Key), "/")}
			acts = append(acts, &action.Action{Key: m.Key, Type: "model", Title: m.Name(), Config: mCfg})
		}
		acts = append(acts, &action.Action{Key: "separator", Type: "separator"})
	}

	sCfg := map[string]string{"source": src.Key, "activity": "sql"}
	sd := "Run ad-hoc SQL queries and DDL"
	acts = append(acts, &action.Action{Key: "sql", Type: "activity", Title: "SQL Playground", Description: sd, Ordinal: 0, Config: sCfg})

	prj := &Project{Key: SourceProjectPrefix + src.Key, Title: src.Title, Description: src.Description, Sources: []string{sourceKey}, Actions: acts}

	return &View{Project: prj, Sources: sources, Schemata: schemata}, nil
}
