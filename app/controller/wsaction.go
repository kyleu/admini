package controller

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/views"
)

func handleAction(req *workspaceRequest, act *action.Action) (string, error) {
	if act == nil {
		return whoops(req, "nil project action", req.Path...)
	}
	msg := "Action [" + act.Key + "] hit!"
	p := append(act.Pkg, req.Path...)
	req.PS.Title = act.Name()
	req.PS.Data = act
	page := &views.TODO{Message: fmt.Sprintf("%v [%v]", msg, strings.Join(req.Path[1:], "/"))}
	return render(req.R, req.W, req.AS, page, req.PS, p...)
}
