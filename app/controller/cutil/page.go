package cutil

import (
	"github.com/gorilla/sessions"
	"github.com/kyleu/admini/app/menu"
)

type PageState struct {
	Menu        menu.Items        `json:"menu,omitempty"`
	Breadcrumbs []string          `json:"breadcrumbs,omitempty"`
	Flashes     []string          `json:"flashes,omitempty"`
	Session     *sessions.Session `json:"-"`
	Icons       []string          `json:"-"`
	RootPath    string            `json:"-"`
	RootTitle   string            `json:"-"`
	SearchPath  string            `json:"-"`
	ProfilePath string            `json:"-"`
	Data        interface{}       `json:"-"`
}

func (p *PageState) AddIcon(icons ...string) {
	for _, test := range icons {
		for _, icon := range p.Icons {
			if icon == test {
				break
			}
		}
		p.Icons = append(p.Icons, test)
	}
}
