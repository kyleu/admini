// Content managed by Project Forge, see [projectforge.md] for details.
package cutil

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"

	"admini.dev/admini/app"
	"admini.dev/admini/app/controller/cmenu"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/lib/theme"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

const (
	DefaultSearchPath  = "/search"
	DefaultProfilePath = "/profile"
	defaultIcon        = "app"
)

var (
	defaultRootTitleAppend = util.GetEnv("app_display_name_append")
	defaultRootTitle       = func() string {
		if tmp := util.GetEnv("app_display_name"); tmp != "" {
			return tmp
		}
		return util.AppName
	}()
)

type PageState struct {
	Title         string            `json:"title,omitempty"`
	Description   string            `json:"description,omitempty"`
	Method        string            `json:"method,omitempty"`
	URI           *fasthttp.URI     `json:"-"`
	Menu          menu.Items        `json:"menu,omitempty"`
	Breadcrumbs   cmenu.Breadcrumbs `json:"breadcrumbs,omitempty"`
	Flashes       []string          `json:"flashes,omitempty"`
	Session       util.ValueMap     `json:"-"`
	Profile       *user.Profile     `json:"profile,omitempty"`
	Accounts      user.Accounts     `json:"accounts,omitempty"`
	Authed        bool              `json:"authed,omitempty"`
	Admin         bool              `json:"admin,omitempty"`
	Icons         []string          `json:"icons,omitempty"`
	RootIcon      string            `json:"rootIcon,omitempty"`
	RootPath      string            `json:"rootPath,omitempty"`
	RootTitle     string            `json:"rootTitle,omitempty"`
	SearchPath    string            `json:"searchPath,omitempty"`
	ProfilePath   string            `json:"profilePath,omitempty"`
	HideMenu      bool              `json:"hideMenu,omitempty"`
	ForceRedirect string            `json:"forceRedirect,omitempty"`
	HeaderContent string            `json:"headerContent,omitempty"`
	Data          any               `json:"data,omitempty"`
	Logger        util.Logger       `json:"-"`
	Context       context.Context   `json:"-"`
	Span          *telemetry.Span   `json:"-"`
	RenderElapsed float64           `json:"renderElapsed,omitempty"`
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

func (p *PageState) User() string {
	if len(p.Accounts) == 0 {
		return "anonymous"
	}
	return p.Accounts[0].Email
}

func (p *PageState) Clean(as *app.State) error {
	if p.Profile != nil && p.Profile.Theme == "" {
		p.Profile.Theme = theme.ThemeDefault.Key
	}
	if p.RootIcon == "" {
		p.RootIcon = defaultIcon
	}
	if p.RootPath == "" {
		p.RootPath = "/"
	}
	if p.RootTitle == "" {
		p.RootTitle = defaultRootTitle
	}
	if defaultRootTitleAppend != "" {
		p.RootTitle += " " + defaultRootTitleAppend
	}
	if p.SearchPath == "" {
		p.SearchPath = DefaultSearchPath
	}
	if p.ProfilePath == "" {
		p.ProfilePath = DefaultProfilePath
	}
	if len(p.Menu) == 0 {
		m, err := cmenu.MenuFor(p.Context, p.Authed, p.Admin, as, p.Logger)
		if err != nil {
			return err
		}
		p.Menu = m
	}
	return nil
}

func (p *PageState) Close() {
	if p.Span != nil {
		p.Span.Complete()
	}
}
