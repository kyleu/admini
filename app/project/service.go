package project

import (
	"context"
	"path/filepath"

	"github.com/pkg/errors"

	"admini.dev/admini/app/lib/filesystem"
	"admini.dev/admini/app/lib/search/result"
	"admini.dev/admini/app/loader"
	"admini.dev/admini/app/source"
	"admini.dev/admini/app/util"
)

type Service struct {
	root    string
	cache   Projects
	files   filesystem.FileLoader
	sources *source.Service
	loaders *loader.Service
}

func NewService(files filesystem.FileLoader, sources *source.Service, ld *loader.Service) *Service {
	return &Service{root: "project", files: files, sources: sources, loaders: ld}
}

func (s *Service) List(_ context.Context, logger util.Logger) (Projects, error) {
	if s.cache == nil {
		err := s.reloadCache(logger)
		if err != nil {
			return nil, err
		}
	}
	return s.cache, nil
}

func (s *Service) Search(ctx context.Context, q string, logger util.Logger) (result.Results, error) {
	ret := result.Results{}
	ps, err := s.List(ctx, logger)
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		if res := result.NewResult("project", p.Key, p.WebPath(), p.Name(), p.IconWithFallback(), p, p, q); len(res.Matches) > 0 {
			ret = append(ret, res)
		}
	}
	return ret, nil
}

func (s *Service) reloadCache(logger util.Logger) error {
	files := s.files.ListJSON(s.root, nil, true, logger)
	ret := make(Projects, 0, len(files))

	for _, key := range files {
		prj, err := s.LoadRequired(key, true, logger)
		if err != nil {
			return errors.Wrapf(err, "unable to load project [%s]", key)
		}
		ret = append(ret, prj)
	}
	ret.Sort()
	s.cache = ret
	return nil
}

func (s *Service) ReloadProject(key string, logger util.Logger) error {
	prj, err := s.LoadRequired(key, true, logger)
	if err != nil {
		return errors.Wrapf(err, "unable to load project [%s]", key)
	}
	s.cache = s.cache.Replace(prj)
	return nil
}

func (s *Service) Load(key string, force bool, logger util.Logger) (*Project, error) {
	if !force {
		if s.cache == nil {
			err := s.reloadCache(logger)
			if err != nil {
				return nil, err
			}
		}
		if curr := s.cache.Get(key); curr != nil {
			return curr, nil
		}
	}

	pf := filepath.Join(s.root, key+".json")

	if !s.files.Exists(pf) {
		return nil, errors.Errorf("no project file found at [%s]", pf)
	}
	ret := &Project{}
	out, err := s.files.ReadFile(pf)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read project [%s]", key)
	}

	err = util.FromJSON(out, ret)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse project")
	}
	ret.Key = key

	s.cache = s.cache.Replace(ret)

	return ret, nil
}

func (s *Service) LoadRequired(key string, force bool, logger util.Logger) (*Project, error) {
	ret, err := s.Load(key, force, logger)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, errors.Errorf("no project found with key [%s]", key)
	}
	ret.Actions.Cleanup()
	return ret, nil
}

func (s *Service) Save(prj *Project, overwrite bool, logger util.Logger) error {
	p := filepath.Join(s.root, prj.Key+".json")
	if !overwrite && s.files.Exists(p) {
		return errors.Errorf("project [%s] already exists", prj.Key)
	}
	j := util.ToJSONBytes(prj, true)
	err := s.files.WriteFile(p, j, filesystem.DefaultMode, overwrite)
	if err != nil {
		return errors.Wrapf(err, "unable to save project [%s]", prj.Key)
	}
	err = s.ReloadProject(prj.Key, logger)
	if err != nil {
		return errors.Wrap(err, "unable to load sources")
	}
	return nil
}

func (s *Service) Delete(key string, logger util.Logger) error {
	p := filepath.Join(s.root, key)
	if !s.files.Exists(p) {
		return errors.Errorf("project [%s] doesn't exist", key)
	}
	err := s.files.RemoveRecursive(p, logger)
	if err != nil {
		return errors.Wrap(err, "unable to remove project files")
	}
	err = s.reloadCache(logger)
	if err != nil {
		return errors.Wrap(err, "unable to load project cache")
	}
	return nil
}

func (s *Service) Clear() {
	s.cache = nil
}
