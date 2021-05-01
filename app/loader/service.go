package loader

import (
	"github.com/kyleu/admini/app/schema"
)

type Service struct {
	loaders map[schema.Origin]Loader
}

func NewService() *Service {
	return &Service{loaders: map[schema.Origin]Loader{}}
}

func (s *Service) Set(o schema.Origin, l Loader) {
	s.loaders[o] = l
}

func (s *Service) Get(origin schema.Origin) Loader {
	return s.loaders[origin]
}
