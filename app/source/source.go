package source

import (
	"encoding/json"
	"sort"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/lib/schema"
	"github.com/kyleu/admini/app/util"
)

type Source struct {
	Key         string          `json:"-"`
	Title       string          `json:"title,omitempty"`
	Icon        string          `json:"icon,omitempty"`
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

func (s *Source) IconWithFallback() string {
	if s.Icon == "" {
		return "app"
	}
	return s.Icon
}

func (s *Source) WebPath() string {
	return "/source/" + s.Key
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

func (s Sources) Keys() []string {
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.Key)
	}
	return ret
}

func (s Sources) Names() []string {
	ret := make([]string, 0, len(s))
	for _, x := range s {
		ret = append(ret, x.Name())
	}
	return ret
}

func (s Sources) GetWithError(key string) (*Source, error) {
	if ret := s.Get(key); ret != nil {
		return ret, nil
	}

	keys := make([]string, 0, len(s))
	for _, x := range s {
		keys = append(keys, x.Key)
	}
	return nil, errors.Errorf("no source [%s] available among candidates [%s]", key, util.StringArrayOxfordComma(keys, "and"))
}

func (s Sources) Sort() {
	sort.Slice(s, func(l int, r int) bool {
		return s[l].Key < s[r].Key
	})
}
