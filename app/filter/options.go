package filter

type Options struct {
	Sort   []string `json:"sort,omitempty"`
	Filter []string `json:"filter,omitempty"`
	Group  []string `json:"group,omitempty"`
	Search string   `json:"search,omitempty"`
	Max    int      `json:"max,omitempty"`
	Params ParamSet `json:"params,omitempty"`
}
