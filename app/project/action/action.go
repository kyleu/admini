package action

import (
	"github.com/kyleu/admini/app/util"
	"sort"
)

type Action struct {
	Key         string            `json:"key"`
	Type        string            `json:"type,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	Ordinal     int               `json:"ordinal,omitempty"`
	Children    Actions           `json:"children,omitempty"` // stored in subdirs
	Pkg         util.Pkg          `json:"-"`
	Config      map[string]string `json:"config,omitempty"`
}

func (a *Action) Name() string {
	if a.Title == "" {
		return a.Key
	}
	return a.Title
}

func (a *Action) Clone(pkg util.Pkg, kids Actions) *Action {
	return &Action{
		Key: a.Key, Type: a.Type, Title: a.Title, Description: a.Description,
		Icon: a.Icon, Ordinal: a.Ordinal, Children: kids, Pkg: pkg, Config: a.Config,
	}
}

type Actions []*Action

func (a Actions) Sort() {
	sort.Slice(a, func(i, j int) bool {
		if a[i].Ordinal == a[j].Ordinal {
			return a[i].Key < a[j].Key
		}
		return a[i].Ordinal < a[j].Ordinal
	})
}

func (a Actions) Get(paths []string) (*Action, []string) {
	if len(paths) == 0 {
		return nil, nil
	}
	curr := a.Find(paths[0])
	if curr == nil {
		return nil, paths
	}
	if len(curr.Children) > 0 {
		x, remaining := curr.Children.Get(paths[1:])
		if x == nil {
			return curr, paths[1:]
		}
		return x, remaining
	}

	return curr, paths
}

func (a Actions) Find(key string) *Action {
	for _, act := range a {
		if act.Key == key {
			return act
		}
	}
	return nil
}

type dto struct {
	Type        string            `json:"type,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Icon        string            `json:"icon,omitempty"`
	Ordinal     int               `json:"ordinal,omitempty"`
	Children    Actions           `json:"-"` // excluded for saving
	Pkg         util.Pkg          `json:"-"`
	Config      map[string]string `json:"config,omitempty"`
}

func newDTO(a *Action) *dto {
	return &dto{
		Type: a.Type, Title: a.Title, Description: a.Description, Icon: a.Icon,
		Ordinal: a.Ordinal, Children: a.Children, Pkg: a.Pkg, Config: a.Config,
	}
}

func (d *dto) ToAction(key string) *Action {
	return &Action{
		Key: key, Type: d.Type, Title: d.Title, Description: d.Description, Icon: d.Icon,
		Ordinal: d.Ordinal, Children: d.Children, Pkg: d.Pkg, Config: d.Config,
	}
}
