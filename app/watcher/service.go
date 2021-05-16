package watcher

import (
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type Service struct {
	watcher *fsnotify.Watcher
	logger *zap.SugaredLogger
}

func NewService(path string, logger *zap.SugaredLogger) (*Service, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}

	l := logger.With(zap.String("service", "watcher"))
	return &Service{watcher: watcher, logger: l}, nil
}

func (s *Service) Close() {
	if s != nil && s.watcher != nil {
		_ = s.watcher.Close()
		s.watcher = nil
	}
}
