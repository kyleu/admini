package source

import (
	"fmt"
	"github.com/kyleu/admini/app/filesystem"
	"github.com/kyleu/admini/app/util"
	"path/filepath"
)

type Service struct {
	root  string
	cache Sources
	files filesystem.FileLoader
}

func NewService(root string, files filesystem.FileLoader) *Service {
	return &Service{root: root, files: files}
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
		return nil, err
	}

	ret := &Source{}
	err = util.FromJSON(out, ret)
	if err != nil {
		return nil, err
	}

	ret.Key = key
	if ret.Title == "" {
		ret.Title = key
	}
	return ret, nil
}
