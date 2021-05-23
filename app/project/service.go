package project

import (
	"path/filepath"

	"go.uber.org/zap"

	"github.com/kyleu/admini/app/project/action"
	"github.com/pkg/errors"

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
	logger  *zap.SugaredLogger
}

func NewService(root string, files filesystem.FileLoader, sources *source.Service, ld *loader.Service, logger *zap.SugaredLogger) *Service {
	log := logger.With(zap.String("service", "project"))
	return &Service{root: root, files: files, sources: sources, loaders: ld, logger: log}
}

func (s *Service) List() (Projects, error) {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Projects, 0, len(dirs))

		for _, dir := range dirs {
			src, err := s.Load(dir, false)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load source [%v]", dir)
			}
			ret = append(ret, src)
		}
		ret.Sort()
		s.cache = ret
	}
	return s.cache, nil
}

func (s *Service) Load(key string, force bool) (*Project, error) {
	if !force {
		if curr := s.cache.Get(key); curr != nil {
			return curr, nil
		}
	}

	dir := filepath.Join(s.root, key)
	pf := filepath.Join(dir, "project.json")

	ret := &Project{}

	if s.files.Exists(pf) {
		out, err := s.files.ReadFile(pf)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read project ["+key+"]")
		}

		err = util.FromJSON(out, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse project")
		}
	}

	ret.Key = key

	actions, err := action.Load(filepath.Join(dir, "actions"), s.files)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load actions")
	}
	ret.Actions = actions
	s.cache.Replace(ret)

	return ret, nil
}

func (s *Service) LoadView(key string) (*View, error) {
	p, err := s.Load(key, false)
	if err != nil {
		return nil, err
	}
	sch, err := s.SchemataFor(key)
	if err != nil {
		return nil, err
	}
	src, err := s.SourcesFor(p)
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
		return errors.Wrapf(err, "unable to save project [%v]", key)
	}
	return nil
}

func (s *Service) SchemataFor(key string) (schema.Schemata, error) {
	p, err := s.Load(key, false)
	if err != nil {
		return nil, errors.Wrapf(err, "can't load project [%v]", key)
	}
	ret := map[string]*schema.Schema{}
	for _, sch := range p.Sources {
		x, err := s.sources.LoadSchema(sch)
		if err != nil {
			return nil, errors.Wrapf(err, "can't load schema [%v] for project [%v]", sch, p.Key)
		}
		ret[sch] = x
	}
	return ret, nil
}

func (s *Service) SourcesFor(p *Project) (source.Sources, error) {
	ret := make(source.Sources, 0, len(p.Sources))
	for _, sch := range p.Sources {
		x, err := s.sources.Load(sch, false)
		if err != nil {
			return nil, errors.Wrapf(err, "can't load source [%v] for project [%v]", sch, p.Key)
		}
		ret = append(ret, x)
	}
	return ret, nil
}
