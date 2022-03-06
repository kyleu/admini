package qualify

import (
	"fmt"
	"strings"

	"admini.dev/app/action"
)

type Qualification struct {
	Action *action.Action `json:"act"`
	Icon   string         `json:"icon,omitempty"`
	Path   []string       `json:"path"`
	Debug  string         `json:"debug,omitempty"`
}

func (q *Qualification) String() string {
	return strings.Join(q.Link(), "/")
}

func (q *Qualification) Help() string {
	return fmt.Sprintf("View %s", q.Action.Title)
}

func (q *Qualification) Link() []string {
	return append(q.Action.Path(), q.Path...)
}

type Qualifications []*Qualification
