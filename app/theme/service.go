package theme

import (
	"github.com/kyleu/admini/app/filesystem"
	"go.uber.org/zap"
)

type Service struct {
	files  filesystem.FileLoader
	cache  Themes
	logger *zap.SugaredLogger
}

func NewService(files filesystem.FileLoader, logger *zap.SugaredLogger) *Service {
	return &Service{files: files, logger: logger}
}

func (s *Service) All() Themes {
	s.loadIfNeeded()
	return s.cache
}
func (s *Service) Clear() {
	s.cache = nil
}

func (s *Service) Get(theme string) *Theme {
	for _, t := range s.All() {
		if t.Key == theme {
			return t
		}
	}
	return ThemeDefault
}

func (s *Service) loadIfNeeded() {
	if s.cache == nil {
		s.cache = Themes{ThemeDefault, ThemeInverse}
		s.cache.Sort()
	}
}

