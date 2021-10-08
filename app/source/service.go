package source

import (
	"path/filepath"

	"github.com/kyleu/admini/app/database"
	"go.uber.org/zap"

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

func NewService(files filesystem.FileLoader, ld *loader.Service, logger *zap.SugaredLogger) *Service {
	log := logger.With(zap.String("service", "source"))
	return &Service{root: "source", schemaCache: map[string]*schema.Schema{}, files: files, loaders: ld, logger: log}
}

func (s *Service) List() (Sources, error) {
	if s.cache == nil {
		err := s.reloadSourceCache()
		if err != nil {
			return nil, err
		}
	}
	return s.cache, nil
}

func (s *Service) NewSource(key string, title string, icon string, description string, t schema.Origin) *Source {
	ret := &Source{Key: key, Title: title, Icon: icon, Description: description, Type: t}
	switch t {
	case schema.OriginMySQL:
		ret.Config = util.ToJSONBytes(&database.MySQLParams{Host: "localhost", Port: 3306, Database: key}, true)
	case schema.OriginPostgres:
		ret.Config = util.ToJSONBytes(&database.PostgresParams{Host: "localhost", Port: 5432, Database: key}, true)
	case schema.OriginSQLite:
		ret.Config = util.ToJSONBytes(&database.SQLiteParams{File: key + ".db"}, true)
	default:
		ret.Config = []byte("{}")
	}
	return ret
}

func (s *Service) Load(key string, force bool) (*Source, error) {
	if !force {
		if s.cache == nil {
			err := s.reloadSourceCache()
			if err != nil {
				return nil, err
			}
		}
		if curr := s.cache.Get(key); curr != nil {
			return curr, nil
		}
	}

	p := filepath.Join(s.root, key, "source.json")

	out, err := s.files.ReadFile(p)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read source [%s]", key)
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
		return errors.Errorf("source [%s] already exists", src.Key)
	}
	f := filepath.Join(p, "source.json")
	j := util.ToJSONBytes(src, true)
	err := s.files.WriteFile(f, j, filesystem.DefaultMode, overwrite)
	if err != nil {
		return errors.Wrapf(err, "unable to save schema [%s]", src.Key)
	}
	err = s.reloadSourceCache()
	if err != nil {
		return errors.Wrap(err, "unable to load schemata")
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
		return errors.Wrap(err, "unable to remove source files")
	}
	delete(s.schemaCache, key)
	err = s.reloadSourceCache()
	if err != nil {
		return errors.Wrap(err, "unable to load sources")
	}
	return nil
}

func (s *Service) reloadSourceCache() error {
	dirs := s.files.ListDirectories(s.root)
	ret := make(Sources, 0, len(dirs))

	for _, dir := range dirs {
		src, err := s.Load(dir, true)
		if err != nil {
			return errors.Wrapf(err, "unable to load source [%s]", dir)
		}
		ret = append(ret, src)
	}
	s.cache = ret
	return nil
}

func (s *Service) GetOverrides(key string) (schema.Overrides, error) {
	op := filepath.Join(s.root, key, "overrides.json")

	if !s.files.Exists(op) {
		return nil, nil
	}

	out, err := s.files.ReadFile(op)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read schema overrides")
	}

	overrides := schema.Overrides{}
	err = util.FromJSON(out, &overrides)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse schema overrides")
	}
	return overrides, nil
}

func (s *Service) SaveOverrides(key string, os schema.Overrides) error {
	op := filepath.Join(s.root, key, "overrides.json")
	os.Sort()
	bytes := util.ToJSONBytes(os, true)
	err := s.files.WriteFile(op, bytes, filesystem.DefaultMode, true)
	if err != nil {
		return errors.Wrap(err, "unable to write schema overrides")
	}
	return nil
}

func (s *Service) loadOverrides(key string, ret *schema.Schema) error {
	overrides, err := s.GetOverrides(key)
	if err != nil {
		return errors.Wrap(err, "unable to read schema overrides")
	}

	for _, o := range overrides {
		err = o.ApplyTo(ret)
		if err != nil {
			return errors.Wrapf(err, "unable to apply override [%s]", o.String())
		}
	}

	return nil
}

func (s *Service) Clear() {
	s.cache = nil
	s.schemaCache = map[string]*schema.Schema{}
}
