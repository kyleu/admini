package project

import (
	"strings"

	"admini.dev/admini/app/util"
	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/source"
)

type View struct {
	Project  *Project        `json:"project"`
	Sources  source.Sources  `json:"sources,omitempty"`
	Schemata schema.Schemata `json:"schemata,omitempty"`
}

func (s *Service) LoadView(key string, logger util.Logger) (*View, error) {
	if strings.HasPrefix(key, SourceProjectPrefix) {
		return s.LoadSourceProject(strings.TrimPrefix(key, SourceProjectPrefix), logger)
	}

	p, err := s.LoadRequired(key, false, logger)
	if err != nil {
		return nil, err
	}
	src, err := s.sourcesFor(p, logger)
	if err != nil {
		return nil, err
	}
	sch, err := s.schemataFor(key, logger)
	if err != nil {
		return nil, err
	}
	return &View{Project: p, Sources: src, Schemata: sch}, nil
}

func (s *Service) schemataFor(key string, logger util.Logger) (schema.Schemata, error) {
	p, err := s.LoadRequired(key, false, logger)
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

func (s *Service) sourcesFor(p *Project, logger util.Logger) (source.Sources, error) {
	ret := make(source.Sources, 0, len(p.Sources))
	for _, sch := range p.Sources {
		x, err := s.sources.Load(sch, false, logger)
		if err != nil {
			return nil, errors.Wrapf(err, "can't load source [%s] for project [%s]", sch, p.Key)
		}
		ret = append(ret, x)
	}
	return ret, nil
}
