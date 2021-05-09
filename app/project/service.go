package project

import (
	"fmt"
	"github.com/kyleu/admini/app/project/action"
	"github.com/pkg/errors"
	"path/filepath"

	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/source"

	"github.com/kyleu/admini/app/loader"

	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
)

type Service struct {
	root    string
	cache   Projects
	files   filesystem.FileLoader
	sources *source.Service
	loaders *loader.Service
}

func NewService(root string, files filesystem.FileLoader, sources *source.Service, ld *loader.Service) *Service {
	return &Service{root: root, files: files, sources: sources, loaders: ld}
}

func (s *Service) List() (Projects, error) {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Projects, 0, len(dirs))

		for _, dir := range dirs {
			src, err := s.Load(dir)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("unable to load source [%v]", dir))
			}
			ret = append(ret, src)
		}
		s.cache = ret
	}
	return s.cache, nil
}

func (s *Service) Load(key string) (*Project, error) {
	curr := s.cache.Get(key)
	if curr != nil {
		return curr, nil
	}

	dir := filepath.Join(s.root, key)
	pf := filepath.Join(dir, "project.json")

	out, err := s.files.ReadFile(pf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read project ["+key+"]")
	}

	ret := &Project{}
	err = util.FromJSON(out, ret)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse project")
	}

	ret.Key = key
	if ret.Title == "" {
		ret.Title = key
	}

	actions, err := action.Load(filepath.Join(dir, "actions"), s.files)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load actions")
	}
	ret.Actions = actions

	return ret, nil
}

func (s *Service) LoadView(key string) (*View, error) {
	p, err := s.Load(key)
	if err != nil {
		return nil, err
	}
	sch, err := s.SchemataFor(key)
	if err != nil {
		return nil, err
	}
	src, err := s.SourcesFor(key)
	if err != nil {
		return nil, err
	}
	return &View{Project: p, Schemata: sch, Sources: src}, nil
}

func (s *Service) SaveProject(key string, prj *Project) error {
	p := filepath.Join(s.root, key, "project.json")
	j := util.ToJSONBytes(prj, true)
	err := s.files.WriteFile(p, j, true)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to save project [%v]", key))
	}
	return nil
}

func (s *Service) SchemataFor(key string) (map[string]*schema.Schema, error) {
	p, err := s.Load(key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't load project [%v]", key))
	}
	ret := map[string]*schema.Schema{}
	for _, sch := range p.Sources {
		x, err := s.sources.SchemaFor(sch)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("can't load schema [%v] for project [%v]", sch, p.Key))
		}
		ret[sch] = x
	}
	return ret, nil
}

func (s *Service) SourcesFor(key string) (map[string]*source.Source, error) {
	p, err := s.Load(key)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't load project [%v]", key))
	}
	ret := map[string]*source.Source{}
	for _, sch := range p.Sources {
		x, err := s.sources.Load(sch)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("can't load source [%v] for project [%v]", sch, p.Key))
		}
		ret[sch] = x
	}
	return ret, nil
}
