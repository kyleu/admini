package source

import (
	"context"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/lib/filesystem"
	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/app/util"
)

func (s *Service) LoadSchema(key string) (*schema.Schema, error) {
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
	if ret != nil {
		err := s.loadOverrides(key, ret)
		if err != nil {
			return nil, errors.Wrap(err, "unable to calculate overrides")
		}

		err = ret.CreateReferences()
		if err != nil {
			return nil, errors.Wrap(err, "unable to calculate references")
		}
	}

	s.schemaCache[key] = ret
	return ret, nil
}

func (s *Service) SaveSchema(key string, sch *schema.Schema) error {
	p := filepath.Join(s.root, key, "schema.json")
	j := util.ToJSONBytes(sch, true)
	err := s.files.WriteFile(p, j, filesystem.DefaultMode, true)
	if err != nil {
		return errors.Wrapf(err, "unable to save schema [%s]", key)
	}
	s.schemaCache[key] = sch
	return nil
}

func (s *Service) SchemaRefresh(ctx context.Context, key string) (*schema.Schema, float64, error) {
	startNanos := time.Now().UnixNano()
	source, err := s.Load(key, false)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "can't load source with key [%s]", key)
	}
	ld, err := s.loaders.Get(source.Type, source.Key, source.Config)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "can't create loader for source [%s]", key)
	}
	if ld == nil {
		return nil, 0, errors.Errorf("no loader defined for type [%s]", source.Type.String())
	}
	sch, err := ld.Schema(ctx)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "can't load schema with key [%s]", key)
	}
	elapsedMillis := float64((time.Now().UnixNano()-startNanos)/int64(time.Microsecond)) / float64(1000)

	err = s.SaveSchema(key, sch)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "can't save source with key [%s]", key)
	}

	err = s.loadOverrides(key, sch)
	if err != nil {
		return nil, 0, errors.Wrap(err, "unable to calculate overrides")
	}
	err = sch.CreateReferences()
	if err != nil {
		return nil, 0, err
	}

	return sch, elapsedMillis, err
}
