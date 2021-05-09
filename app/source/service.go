package source

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"time"

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
}

func NewService(root string, files filesystem.FileLoader, ld *loader.Service) *Service {
	return &Service{root: root, schemaCache: map[string]*schema.Schema{}, files: files, loaders: ld}
}

func (s *Service) List() (Sources, error) {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Sources, 0, len(dirs))

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

func (s *Service) Load(key string) (*Source, error) {
	curr := s.cache.Get(key)
	if curr != nil {
		return curr, nil
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
	if ret.Title == "" {
		ret.Title = key
	}
	return ret, nil
}

func (s *Service) SchemaFor(key string) (*schema.Schema, error) {
	curr, ok := s.schemaCache[key]
	if ok {
		return curr, nil
	}
	var ret *schema.Schema
	p := filepath.Join(s.root, key, "schema.json")

	if s.files.Exists(p) {
		out, err := s.files.ReadFile(p)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read schema")
		}

		ret = &schema.Schema{}
		err = util.FromJSON(out, ret)
		if err != nil {
			return nil, err
		}
	}
	s.schemaCache[key] = ret
	return ret, nil
}

func (s *Service) SchemaRefresh(key string) (*schema.Schema, float64, error) {
	startNanos := time.Now().UnixNano()
	source, err := s.Load(key)
	if err != nil {
		return nil, 0, errors.Wrap(err, fmt.Sprintf("can't load source with key [%s]", key))
	}
	ld := s.loaders.Get(source.Type)
	if ld == nil {
		return nil, 0, errors.New(fmt.Sprintf("no loader defined for type [" + source.Type.String() + "]"))
	}
	sch, err := ld.Schema(source.Key, source.Config)
	if err != nil {
		return nil, 0, errors.Wrap(err, fmt.Sprintf("can't load schema with key [%s]", key))
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

	err = s.SaveSchema(key, sch)
	if err != nil {
		return nil, 0, errors.Wrap(err, fmt.Sprintf("can't save source with key [%s]", key))
	}

	return sch, elapsedMillis, err
}

func (s *Service) SaveSchema(key string, sch *schema.Schema) error {
	p := filepath.Join(s.root, key, "schema.json")
	j := util.ToJSONBytes(sch, true)
	err := s.files.WriteFile(p, j, true)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to save schema [%v]", key))
	}
	s.schemaCache[key] = sch
	return nil
}
