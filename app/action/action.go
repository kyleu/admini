package action

import (
	"fmt"

	"admini.dev/admini/app/util"
)

const defaultIcon = "app"

var RootAction = &Action{TypeKey: TypeFolder.Key, Config: util.ValueMap{}}

type Action struct {
	Key         string        `json:"key,omitempty"`
	TypeKey     string        `json:"type,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Icon        string        `json:"icon,omitempty"`
	Children    Actions       `json:"children,omitempty"`
	Pkg         util.Pkg      `json:"-"` // injected
	Config      util.ValueMap `json:"config,omitempty"`
}

func (a *Action) Name() string {
	if a.Title == "" {
		return a.Key
	}
	return a.Title
}

func (a *Action) Type() *Type {
	ret, _ := TypeFromString(a.TypeKey)
	return ret
}

func (a *Action) ConfigString() string {
	ret := a.Type().ConfigString(a.Config)
	if len(a.Children) == 0 {
		return ret
	}
	kids := util.StringPlural(len(a.Children), "child")
	if ret == "" {
		return kids
	}
	return fmt.Sprintf("%s, %s", ret, kids)
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
	if a.Type().Icon != "" {
		return a.Type().Icon
	}
	return defaultIcon
}

func (a *Action) Clone(pkg util.Pkg, kids Actions) *Action {
	return &Action{
		Key: a.Key, TypeKey: a.TypeKey, Title: a.Title, Description: a.Description,
		Icon: a.Icon, Children: kids, Pkg: pkg, Config: a.Config,
	}
}

func (a *Action) String() string {
	return fmt.Sprintf("workspace:%s:%s", a.TypeKey, a.Key)
}
