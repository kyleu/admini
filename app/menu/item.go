package menu

type Item struct {
	Key         string `json:"key"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Route       string `json:"route,omitempty"`
	Children    Items  `json:"children,omitempty"`
}

func (i *Item) AddChild(child *Item) {
	i.Children = append(i.Children, child)
}

type Items []*Item
