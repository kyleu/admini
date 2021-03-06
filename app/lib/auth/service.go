// Content managed by Project Forge, see [projectforge.md] for details.
package auth

import (
	"admini.dev/admini/app/util"
)

type Service struct {
	baseURL   string
	providers Providers
}

func NewService(baseURL string, logger util.Logger) *Service {
	ret := &Service{baseURL: baseURL}
	_ = ret.load(logger)
	return ret
}

func (s *Service) LoginURL() string {
	if len(s.providers) == 1 {
		return "/auth/" + s.providers[0].ID
	}
	return defaultProfilePath
}
