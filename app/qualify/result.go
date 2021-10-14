package qualify

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
)

type Result struct {
	Action *action.Action `json:"act"`
	Icon   string         `json:"icon,omitempty"`
	Path   []string       `json:"path"`
	Model  *model.Model   `json:"model,omitempty"`
	Data   []interface{}  `json:"data,omitempty"`
	Debug  string         `json:"debug,omitempty"`
}

func (r *Result) ToResult() *result.Result {
	return &result.Result{
		Title:  r.Action.Title,
		Count:  len(r.Data),
		Query:  "todo",
		Fields: r.Model.Fields,
		Data:   [][]interface{}{r.Data},
		Timing: &result.Timing{},
	}
}

func (r *Result) String() string {
	link := strings.Join(append(r.Action.Path(), r.Path...), "/")
	if r.Debug == "" {
		return link
	}
	return fmt.Sprintf("%s (%s)", link, r.Debug)
}

type Results []*Result
