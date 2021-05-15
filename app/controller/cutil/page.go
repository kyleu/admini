package cutil

import (
	"net/url"

	"go.uber.org/zap"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/util"
)

type PageState struct {
	Title       string
	Description string
	Method      string
	URL         *url.URL
	Menu        menu.Items
	Breadcrumbs []string
	Flashes     []string
	Session     *sessions.Session
	Icons       []string
	RootPath    string
	RootTitle   string
	SearchPath  string
	ProfilePath string
	Data        interface{}
	Logger      *zap.SugaredLogger
}

func (p *PageState) AddIcon(n string) {
	for _, icon := range p.Icons {
		if icon == n {
			return
		}
	}
	p.Icons = append(p.Icons, n)
}

func (p *PageState) TitleString() string {
	if p.Title == "" {
		return util.AppName
	}
	return p.Title + " - " + util.AppName
}
