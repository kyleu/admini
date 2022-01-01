package project

import (
	"strings"

	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/app/source"
	"github.com/pkg/errors"
)

type View struct {
	Project  *Project        `json:"project"`
	Sources  source.Sources  `json:"sources,omitempty"`
	Schemata schema.Schemata `json:"schemata,omitempty"`
}

func (s *Service) LoadView(key string) (*View, error) {
	if strings.HasPrefix(key, SourceProjectPrefix) {
		return s.LoadSourceProject(strings.TrimPrefix(key, SourceProjectPrefix))
	}

	p, err := s.LoadRequired(key, false)
	if err != nil {
		return nil, err
	}
	src, err := s.sourcesFor(p)
	if err != nil {
		return nil, err
	}
	sch, err := s.schemataFor(key)
	if err != nil {
		return nil, err
	}
	return &View{Project: p, Sources: src, Schemata: sch}, nil
}

func (s *Service) schemataFor(key string) (schema.Schemata, error) {
	p, err := s.LoadRequired(key, false)
	if err != nil {
		return nil, errors.Wrapf(err, "can't load project [%s]", key)
	}
	ret := map[string]*schema.Schema{}
	for _, sch := range p.Sources {
		x, err := s.sources.LoadSchema(sch)
		if err != nil {
			return nil, errors.Wrapf(err, "can't load schema [%s] for project [%s]", sch, p.Key)
		}
		ret[sch] = x
	}
	return ret, nil
}

func (s *Service) sourcesFor(p *Project) (source.Sources, error) {
	ret := make(source.Sources, 0, len(p.Sources))
	for _, sch := range p.Sources {
		x, err := s.sources.Load(sch, false)
		if err != nil {
			return nil, errors.Wrapf(err, "can't load source [%s] for project [%s]", sch, p.Key)
		}
		ret = append(ret, x)
	}
	return ret, nil
}
