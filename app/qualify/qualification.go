package qualify

import (
	"fmt"
	"github.com/kyleu/admini/app/action"
	"strings"
)

type Qualification struct {
	Action *action.Action `json:"act"`
	Icon   string         `json:"icon,omitempty"`
	Path   []string       `json:"path"`
	Debug  string         `json:"debug,omitempty"`
}

func (r *Qualification) String() string {
	if r.Debug == "" {
		return strings.Join(r.Link(), "/")
	}
	return fmt.Sprintf("%s (%s)", strings.Join(r.Link(), "/"), r.Debug)
}

func (r *Qualification) Link() []string {
	return append(r.Action.Path(), r.Path...)
}

type Qualifications []*Qualification
