package watcher

import (
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"

	"admini.dev/admini/app/util"
)

type Service struct {
	watcher *fsnotify.Watcher
	logger  util.Logger
}

func NewService(path string, logger util.Logger) (*Service, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logger.Infof("event [%s]", event.String())
				if event.Op&fsnotify.Write == fsnotify.Write {
					logger.Infof("modified file [%s]", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Errorf("error: %+v", err)
			}
		}
	}()

	l := logger.With(zap.String("service", "watcher"))
	return &Service{watcher: watcher, logger: l}, nil
}

func (s *Service) Close() {
	if s != nil && s.watcher != nil {
		_ = s.watcher.Close()
		s.watcher = nil
	}
}
