package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/markbates/goth"
	"github.com/pkg/errors"
)

type Provider struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Key      string `json:"-"`
	Secret   string `json:"-"`
	Callback string `json:"callback"`
	goth     goth.Provider
}

type Providers []*Provider

func (p Providers) toGoth() ([]goth.Provider, error) {
	ret := make([]goth.Provider, 0, len(p))
	for _, x := range p {
		ret = append(ret, x.goth)
	}
	return ret, nil
}

func (p Providers) Get(id string) *Provider {
	for _, x := range p {
		if x.ID == id {
			return x
		}
	}
	return nil
}

func (p Providers) Contains(id string) bool {
	return p.Get(id) != nil
}

func (p Providers) IDs() []string {
	ret := make([]string, 0, len(p))
	for _, x := range p {
		ret = append(ret, x.ID)
	}
	return ret
}

func (p Providers) Titles() []string {
	ret := make([]string, 0, len(p))
	for _, x := range p {
		ret = append(ret, x.Title)
	}
	return ret
}

func (s *Service) Providers() (Providers, error) {
	if s.providers == nil {
		err := s.load()
		if err != nil {
			return nil, err
		}
	}
	return s.providers, nil
}

func (s *Service) load() error {
	if s.providers != nil {
		return errors.New("called [load] twice")
	}
	if s.baseURL == "" {
		s.baseURL = "http://localhost:14000"
	}
	s.baseURL = strings.TrimSuffix(s.baseURL, "/")

	initAvailable()

	ret := Providers{}
	for _, k := range AvailableProviderKeys {
		u := strings.ToUpper(k)
		envKey := os.Getenv(u + "_KEY")
		if envKey != "" {
			envSecret := os.Getenv(u + "_SECRET")
			cb := fmt.Sprintf("%s/auth/%s/callback", s.baseURL, k)
			gothPrv, err := toGoth(k, envKey, envSecret, cb)
			if err != nil {
				return err
			}
			ret = append(ret, &Provider{ID: k, Title: AvailableProviderNames[k], Key: envKey, Secret: envSecret, Callback: cb, goth: gothPrv})
		}
	}

	gps, err := ret.toGoth()
	if err != nil {
		return err
	}
	goth.UseProviders(gps...)
	s.providers = ret

	if len(ret) == 0 {
		s.logger.Debug("authentication disabled, no providers configured in environment")
	} else {
		s.logger.Debugf("authentication enabled for [%s]", strings.Join(ret.Titles(), ", "))
	}

	return nil
}
