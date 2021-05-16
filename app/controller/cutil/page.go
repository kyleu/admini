package cutil

import (
	"net/url"

	"go.uber.org/zap"

	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/util"
)

type PageState struct {
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Method        string             `json:"method"`
	URL           *url.URL           `json:"-"`
	Menu          menu.Items         `json:"menu"`
	Breadcrumbs   []string           `json:"breadcrumbs"`
	Flashes       []string           `json:"flashes"`
	Session       *sessions.Session  `json:"-"`
	Icons         []string           `json:"icons"`
	RootPath      string             `json:"rootPath"`
	RootTitle     string             `json:"rootTitle"`
	SearchPath    string             `json:"searchPath"`
	ProfilePath   string             `json:"profilePath"`
	Data          interface{}        `json:"data"`
	Logger        *zap.SugaredLogger `json:"-"`
	RenderElapsed float64            `json:"renderElapsed"`
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
