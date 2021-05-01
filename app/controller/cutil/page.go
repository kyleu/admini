package cutil

import (
	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/util"
	"net/url"
)

type PageState struct {
	Title       string
	Description string
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
