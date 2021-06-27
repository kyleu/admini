package qualify

import (
	"fmt"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/model"
	"strings"
)

type Result struct {
	Action *action.Action `json:"act"`
	Icon   string         `json:"icon,omitempty"`
	Path   []string       `json:"path"`
	Model  *model.Model   `json:"-"`
	Data   []interface{}  `json:"-"`
	Debug  string         `json:"debug,omitempty"`
}

func (r *Result) ToWorkspaceRequest(in *cutil.WorkspaceRequest) *cutil.WorkspaceRequest {
	ret := in.Clone()
	ret.Path = r.Path
	return ret
}

func (r *Result) String() string {
	link := strings.Join(append(r.Action.Path(), r.Path...), "/")
	if r.Debug == "" {
		return link
	}
	return fmt.Sprintf("%s (%s)", link, r.Debug)
}

type Results []*Result
