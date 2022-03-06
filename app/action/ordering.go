package action

import (
	"admini.dev/app/util"
	"github.com/pkg/errors"
)

type Ordering struct {
	Key          string    `json:"k"`
	Title        string    `json:"t"`
	OriginalPath string    `json:"p"`
	Children     Orderings `json:"c,omitempty"`
}

type Orderings []*Ordering

func (a Orderings) Get(paths []string) (*Ordering, []string) {
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

	return curr, paths[1:]
}

func (a Orderings) Find(key string) *Ordering {
	for _, act := range a {
		if act.Key == key {
			return act
		}
	}
	return nil
}

func ReorderActions(acts Actions, orderings Orderings) (Actions, error) {
	ret := Actions{}

	for _, o := range orderings {
		act, err := forOrdering(acts, o, util.Pkg{})
		if err != nil {
			return nil, err
		}
		ret = append(ret, act)
	}

	ret.CleanKeys()
	return ret, nil
}

func forOrdering(acts Actions, o *Ordering, pkg util.Pkg) (*Action, error) {
	var act *Action
	p := util.StringSplitAndTrim(o.OriginalPath, "/")
	if o.Key == "_new" {
		var err error
		if len(p) < 1 {
			return nil, errors.New("attempted to create new action with no arguments")
		}
		t, err := TypeFromString(p[0])
		if err != nil {
			return nil, err
		}
		act, err = newAction(p[1:], o.Title, t, pkg)
		if err != nil {
			return nil, errors.Wrapf(err, "can't parse new action from [%s]", o.OriginalPath)
		}
	} else {
		x := append(append([]string{}, p...), o.Key)
		act, _ = acts.Get(x)
		if act == nil {
			return nil, errors.Errorf("no original action available at path [%s]", o.OriginalPath)
		}
	}
	kids := make(Actions, 0, len(o.Children))
	for _, x := range o.Children {
		kid, err := forOrdering(acts, x, pkg.Push(o.Key))
		if err != nil {
			return nil, err
		}
		kids = append(kids, kid)
	}
	kids.CleanKeys()
	cl := act.Clone(pkg, kids)
	return cl, nil
}
