package controller

import (
	"fmt"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/views"
	"strings"
)

func handleAction(req *workspaceRequest, act *action.Action) (string, error) {
	if act == nil {
		return whoops(req, "nil project action")
	}
	msg := "Action [" + act.Key + "] hit!"
	p := append(act.Pkg, req.Path...)
	req.PS.Data = act
	page := &views.TODO{Message: fmt.Sprintf("%v [%v]", msg, strings.Join(p, "/"))}
	return render(req.R, req.W, req.AS, page, req.PS, p...)
}
