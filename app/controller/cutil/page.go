package cutil

import (
	"fmt"

	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/user"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"github.com/go-gem/sessions"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/util"
)

type PageState struct {
	Title         string             `json:"title"`
	Description   string             `json:"description"`
	Method        string             `json:"method"`
	URI           *fasthttp.URI      `json:"-"`
	Menu          menu.Items         `json:"menu"`
	Breadcrumbs   Breadcrumbs        `json:"breadcrumbs"`
	Flashes       []string           `json:"flashes"`
	Session       *sessions.Session  `json:"-"`
	Profile       *user.Profile      `json:"profile"`
	Auth          auth.Sessions      `json:"auth"`
	Icons         []string           `json:"icons"`
	RootIcon      string             `json:"rootIcon"`
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
	return fmt.Sprintf("%s - %s", p.Title, util.AppName)
}
