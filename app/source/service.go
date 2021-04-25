package source

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/kyleu/admini/app/loader"

	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

type Service struct {
	root    string
	cache   Sources
	files   filesystem.FileLoader
	loaders *loader.Service
}

func NewService(root string, files filesystem.FileLoader, ld *loader.Service) *Service {
	return &Service{root: root, files: files, loaders: ld}
}

func (s *Service) List() (Sources, error) {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Sources, 0, len(dirs))

		for _, dir := range dirs {
			src, err := s.Load(dir)
			if err != nil {
				return nil, fmt.Errorf("unable to load source [%v]: %w", dir, err)
			}
			ret = append(ret, src)
		}
		s.cache = ret
	}
	return s.cache, nil
}

func (s *Service) Load(key string) (*Source, error) {
	p := filepath.Join(s.root, key, "source.json")

	out, err := s.files.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("unable to read source ["+key+"]: %w", err)
	}

	ret := &Source{}
	err = util.FromJSON(out, ret)
	if err != nil {
		return nil, fmt.Errorf("unable : %w", err)
	}

	ret.Key = key
	if ret.Title == "" {
		ret.Title = key
	}
	return ret, nil
}

func (s *Service) SchemaFor(key string) (*schema.Schema, error) {
	var ret *schema.Schema
	p := filepath.Join(s.root, key, "schema.json")

	if s.files.Exists(p) {
		out, err := s.files.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("unable to read schema: %w", err)
		}

		ret = &schema.Schema{}
		err = util.FromJSON(out, ret)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (s *Service) SchemaRefresh(key string) (*schema.Schema, float64, error) {
	startNanos := time.Now().UnixNano()
	source, err := s.Load(key)
	if err != nil {
		return nil, 0, fmt.Errorf("can't load source with key [%s]: %w", key, err)
	}
	ld := s.loaders.Get(source.Type)
	if ld == nil {
		return nil, 0, fmt.Errorf("no loader defined for type [" + source.Type.String() + "]")
	}
	sch, err := ld.GetSchema(source.Config)
	if err != nil {
		return nil, 0, fmt.Errorf("can't load schema with key [%s]: %w", key, err)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

	err = s.SaveSchema(key, sch)
	if err != nil {
		return nil, 0, fmt.Errorf("can't save source with key [%s]: %w", key, err)
	}

	return sch, elapsedMillis, err
}

func (s *Service) SaveSchema(key string, sch *schema.Schema) error {
	p := filepath.Join(s.root, key, "schema.json")
	j := util.ToJSONBytes(sch, true)
	err := s.files.WriteFile(p, j, true)
	if err != nil {
		return fmt.Errorf("unable to save schema [%v]: %w", key, err)
	}
	return nil
}
