package project

import (
	"path/filepath"

	"github.com/kyleu/admini/app/action"

	"go.uber.org/zap"

	"github.com/pkg/errors"

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
		err := s.reloadCache()
		if err != nil {
			return nil, err
		}
	}
	return s.cache, nil
}

func (s *Service) reloadCache() error {
	dirs := s.files.ListDirectories(s.root)
	ret := make(Projects, 0, len(dirs))

	for _, dir := range dirs {
		prj, err := s.LoadRequired(dir, true)
		if err != nil {
			return errors.Wrapf(err, "unable to load project [%s]", dir)
		}
		ret = append(ret, prj)
	}
	ret.Sort()
	s.cache = ret
	return nil
}

func (s *Service) Load(key string, force bool) (*Project, error) {
	if !force {
		if s.cache == nil {
			err := s.reloadCache()
			if err != nil {
				return nil, err
			}
		}
		if curr := s.cache.Get(key); curr != nil {
			return curr, nil
		}
	}

	dir := filepath.Join(s.root, key)
	if !s.files.Exists(dir) {
		return nil, nil
	}
	pf := filepath.Join(dir, "project.json")

	ret := &Project{}
	if s.files.Exists(pf) {
		ret = &Project{}
		out, err := s.files.ReadFile(pf)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to read project [%s]", key)
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

	s.cache = s.cache.Replace(ret)

	return ret, nil
}

func (s *Service) LoadRequired(key string, force bool) (*Project, error) {
	ret, err := s.Load(key, force)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, errors.Errorf("no project found with key [%s]", key)
	}
	return ret, nil
}

func (s *Service) Save(prj *Project, overwrite bool) error {
	p := filepath.Join(s.root, prj.Key)
	if !overwrite && s.files.Exists(p) {
		return errors.Errorf("project [%s] already exists", prj.Key)
	}
	f := filepath.Join(p, "project.json")
	j := util.ToJSONBytes(prj, true)
	err := s.files.WriteFile(f, j, overwrite)
	if err != nil {
		return errors.Wrapf(err, "unable to save project [%s]", prj.Key)
	}
	err = s.reloadCache()
	if err != nil {
		return errors.Wrap(err, "unable to load sources")
	}
	return nil
}

func (s *Service) Delete(key string) error {
	p := filepath.Join(s.root, key)
	if !s.files.Exists(p) {
		return errors.Errorf("source [%s] doesn't exist", key)
	}
	err := s.files.RemoveRecursive(p)
	if err != nil {
		return errors.Wrap(err, "unable to remove project files")
	}
	err = s.reloadCache()
	if err != nil {
		return errors.Wrap(err, "unable to load project cache")
	}
	return nil
}
