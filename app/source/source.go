package source

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"

	"github.com/kyleu/admini/app/schema"
)

type Source struct {
	Key         string          `json:"-"`
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Paths       []string        `json:"paths,omitempty"`
	Type        schema.Origin   `json:"type,omitempty"`
	Config      json.RawMessage `json:"config,omitempty"`
}

func (s *Source) Name() string {
	if s.Title == "" {
		return s.Key
	}
	return s.Title
}

type Sources []*Source

func (s Sources) Get(key string) *Source {
	for _, x := range s {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (s Sources) GetWithError(key string) (*Source, error) {
	ret := s.Get(key)
	if ret != nil {
		return ret, nil
	}

	keys := make([]string, 0, len(s))
	for _, x := range s {
		keys = append(keys, x.Key)
	}
	return nil, errors.Errorf("no source [%v] available among candidates [%v]", key, strings.Join(keys, ", "))
}
