package loader

import (
	"github.com/kyleu/admini/app/schema"
	"github.com/pkg/errors"
)

type loaderFn = func(key string, cfg []byte) (Loader, error)

type Service struct {
	loaders map[schema.Origin]loaderFn
	cache   map[string]Loader
}

func NewService() *Service {
	return &Service{
		loaders: map[schema.Origin]loaderFn{},
		cache:   map[string]Loader{},
	}
}

func (s *Service) Set(o schema.Origin, f loaderFn) {
	s.loaders[o] = f
}

func (s *Service) Get(origin schema.Origin, key string, cfg []byte) (Loader, error) {
	check := origin.Key + "-" + key
	curr, ok := s.cache[check]
	if ok {
		return curr, nil
	}

	o := s.loaders[origin]
	l, err := o(key, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can't create loader")
	}
	s.cache[check] = l
	return l, nil
}

func (s *Service) Clear() {
	s.cache = map[string]Loader{}
}
