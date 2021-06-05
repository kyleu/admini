package action

import (
	"github.com/kyleu/admini/app/util"
)

type Action struct {
	Key         string        `json:"key"`
	Type        Type          `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Icon        string        `json:"icon,omitempty"`
	Ordinal     int           `json:"ordinal,omitempty"`
	Children    Actions       `json:"children,omitempty"` // stored in subdirs
	Pkg         util.Pkg      `json:"-"`
	Config      util.ValueMap `json:"config,omitempty"`
}

func (a *Action) Name() string {
	if a.Title == "" {
		return a.Key
	}
	return a.Title
}

func (a *Action) Path() []string {
	if a.Key == "" {
		return a.Pkg
	}
	return append(a.Pkg, a.Key)
}

func (a *Action) Size() int {
	ret := 1
	for _, x := range a.Children {
		ret += x.Size()
	}
	return ret
}

func (a *Action) IconWithFallback() string {
	if a.Icon != "" {
		return a.Icon
	}
	if a.Type.Icon != "" {
		return a.Type.Icon
	}
	return "app"
}

func (a *Action) Clone(pkg util.Pkg, kids Actions) *Action {
	return &Action{
		Key: a.Key, Type: a.Type, Title: a.Title, Description: a.Description,
		Icon: a.Icon, Ordinal: a.Ordinal, Children: kids, Pkg: pkg, Config: a.Config,
	}
}

type dto struct {
	Type        string        `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Icon        string        `json:"icon,omitempty"`
	Ordinal     int           `json:"ordinal,omitempty"`
	Children    Actions       `json:"-"` // excluded for saving
	Pkg         util.Pkg      `json:"-"`
	Config      util.ValueMap `json:"config,omitempty"`
}

func newDTO(a *Action) *dto {
	return &dto{
		Type: a.Type.Key, Title: a.Title, Description: a.Description, Icon: a.Icon,
		Ordinal: a.Ordinal, Children: a.Children, Pkg: a.Pkg, Config: a.Config,
	}
}

func (d *dto) ToAction(key string) *Action {
	t, _ := TypeFromString(d.Type)
	return &Action{
		Key: key, Type: t, Title: d.Title, Description: d.Description, Icon: d.Icon,
		Ordinal: d.Ordinal, Children: d.Children, Pkg: d.Pkg, Config: d.Config,
	}
}
