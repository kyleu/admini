package source

import (
	"go.uber.org/zap"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/loader"

	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Service struct {
	root        string
	cache       Sources
	schemaCache map[string]*schema.Schema
	files       filesystem.FileLoader
	loaders     *loader.Service
	logger      *zap.SugaredLogger
}

func NewService(root string, files filesystem.FileLoader, ld *loader.Service, logger *zap.SugaredLogger) *Service {
	log := logger.With(zap.String("service", "source"))
	return &Service{root: root, schemaCache: map[string]*schema.Schema{}, files: files, loaders: ld, logger: log}
}

func (s *Service) List() (Sources, error) {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Sources, 0, len(dirs))

		for _, dir := range dirs {
			src, err := s.Load(dir, false)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to load source [%v]", dir)
			}
			ret = append(ret, src)
		}
		s.cache = ret
	}
	return s.cache, nil
}

func (s *Service) Load(key string, force bool) (*Source, error) {
	if !force {
		if curr := s.cache.Get(key); curr != nil {
			return curr, nil
		}
	}

	p := filepath.Join(s.root, key, "source.json")

	out, err := s.files.ReadFile(p)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read source ["+key+"]")
	}

	ret := &Source{}
	err = util.FromJSON(out, ret)
	if err != nil {
		return nil, errors.Wrap(err, "unable to deserialize source")
	}

	ret.Key = key
	return ret, nil
}

func (s *Service) Save(src *Source, overwrite bool) error {
	p := filepath.Join(s.root, src.Key)
	if !overwrite && s.files.Exists(p) {
		return errors.Errorf("source [%v] already exists", src.Key)
	}
	f := filepath.Join(p, "source.json")
	j := util.ToJSONBytes(src, true)
	err := s.files.WriteFile(f, j, true)
	if err != nil {
		return errors.Wrapf(err, "unable to save schema [%v]", src.Key)
	}
	s.cache.Add(src)
	return nil
}
