package source

import (
	"github.com/kyleu/admini/app/filesystem"
)

type Service struct {
	root  string
	cache Sources
	files filesystem.FileLoader
}

func NewService(root string, files filesystem.FileLoader) *Service {
	return &Service{root: root, files: files}
}

func (s *Service) List() Sources {
	if s.cache == nil {
		dirs := s.files.ListDirectories(s.root)
		ret := make(Sources, 0, len(dirs))

		for _, dir := range dirs {
			ret = append(ret, &Source{
				Key:   dir,
				Title: dir,
			})
		}
		s.cache = ret
	}
	return s.cache
}
