package source

import (
	"encoding/json"
	"github.com/pkg/errors"
	"sort"
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

func (s Sources) Add(src *Source) {
	for idx, x := range s {
		if x.Key == src.Key {
			s[idx] = src
			return
		}
	}
	s = append(s, src)
	s.Sort()
	return
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

func (s Sources) Sort() {
	sort.Slice(s, func(l int, r int) bool {
		return s[l].Key < s[r].Key
	})
}
