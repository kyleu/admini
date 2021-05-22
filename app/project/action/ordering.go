package action

type Ordering struct {
	Key          string    `json:"k"`
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
