package action

import (
	"github.com/kyleu/admini/app/util"
	"sort"
)

type Action struct {
	Key         string   `json:"key"`
	Type        Type     `json:"type"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Icon        string   `json:"icon,omitempty"`
	Ordinal     int      `json:"ordinal"`
	Children    Actions  `json:"-"`
	Pkg         util.Pkg `json:"-"`
}

func (a Action) TitleString() string {
	if a.Title == "" {
		return a.Key
	}
	return a.Title
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
	if len(paths) == 1 {
		return curr, paths[1:]
	}
	if len(curr.Children) > 0 {
		x, remaining := curr.Children.Get(paths[1:])
		if x == nil {
			return curr, paths[1:]
		}
		return x, remaining
	}

	return nil, paths
}

func (a Actions) Find(key string) *Action {
	for _, act := range a {
		if act.Key == key {
			return act
		}
	}
	return nil
}
