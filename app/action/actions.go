package action

import (
	"fmt"
	"strings"
)

type Actions []*Action

func (a Actions) Size() int {
	ret := 0
	for _, x := range a {
		ret += x.Size()
	}
	return ret
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

	return curr, paths[1:]
}

func (a Actions) Find(key string) *Action {
	for _, act := range a {
		if act.Key == key {
			return act
		}
	}
	return nil
}

func (a Actions) CleanKeys() {
	for _, act := range a {
		if !strings.HasPrefix(act.Key, "__") {
			continue
		}
		proposed := strings.TrimPrefix(act.Key, "__")
		match := a.Find(proposed)
		idx := 1
		for match != nil {
			idx++
			match = a.Find(fmt.Sprintf("%s-%d", proposed, idx))
		}
		if idx == 1 {
			act.Key = proposed
		} else {
			act.Key = fmt.Sprintf("%s-%d", proposed, idx)
		}
	}
}

func (a Actions) Cleanup() {
	a.cleanup()
}

func (a Actions) cleanup(path ...string) {
	for _, kid := range a {
		kid.Pkg = path
		kid.Children.cleanup(append(append([]string{}, path...), kid.Key)...)
	}
}

func (a Actions) Clone() Actions {
	ret := make(Actions, 0, len(a))
	for _, x := range a {
		ret = append(ret, x.Clone(x.Pkg, x.Children.Clone()))
	}
	return ret
}
