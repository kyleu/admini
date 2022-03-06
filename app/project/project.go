package project

import (
	"fmt"
	"sort"

	"admini.dev/app/action"
)

type Project struct {
	Key         string         `json:"-"`
	Title       string         `json:"title,omitempty"`
	Icon        string         `json:"icon,omitempty"`
	Description string         `json:"description,omitempty"`
	Sources     []string       `json:"sources,omitempty"`
	Actions     action.Actions `json:"actions"`
}

func (p *Project) Name() string {
	if p.Title == "" {
		return p.Key
	}
	return p.Title
}

func (p *Project) IconWithFallback() string {
	if p.Icon == "" {
		return "app"
	}
	return p.Icon
}

func (p *Project) WebPath() string {
	return fmt.Sprintf("/project/%s", p.Key)
}

type Projects []*Project

func (p Projects) Get(key string) *Project {
	for _, x := range p {
		if x.Key == key {
			return x
		}
	}
	return nil
}

func (p Projects) Replace(n *Project) Projects {
	for idx, x := range p {
		if x.Key == n.Key {
			p[idx] = n
			return p
		}
	}
	return append(p, n)
}

func (p Projects) Sort() {
	sort.Slice(p, func(l int, r int) bool {
		return p[l].Key < p[r].Key
	})
}
