package workspace

import (
	"net/http"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/views/vaction"
	"github.com/pkg/errors"
)

func sourceActivity(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	switch act.Config["activity"] {
	case "sql":
		return sourceActivitySQL(req, act)
	case "":
		return ErrResult(req, act, errors.New("must provide activity in action config"))
	default:
		return ErrResult(req, act, errors.New("invalid activity ["+act.Config["activity"]+"] in action config"))
	}
}

func sourceActivitySQL(req *cutil.WorkspaceRequest, act *action.Action) (*Result, error) {
	srcKey := act.Config["source"]
	if srcKey == "" {
		return ErrResult(req, act, errors.New("must provide source in action config"))
	}
	sql := act.Config["query"]
	if req.R.Method == http.MethodPost {
		_ = req.R.ParseForm()
		s := req.R.Form.Get("sql")
		if s != "" {
			sql = s
		}
	}

	if sql == "" {
		page := &vaction.ResultActivitySQL{Req: req, Act: act, SQL: sql, Res: nil}
		return NewResult("", nil, req, act, act, page), nil
	}
	src := req.Sources.Get(srcKey)
	if src == nil {
		return ErrResult(req, act, errors.New("source ["+srcKey+"] is not in this project's configuration"))
	}
	ld, err := req.AS.Loaders.Get(src.Type, src.Key, src.Config)
	if err != nil {
		return ErrResult(req, act, errors.New("unable to create loader"))
	}

	r, err := ld.Query(sql)
	if err != nil {
		return ErrResult(req, act, errors.New("unable to execute query"))
	}
	page := &vaction.ResultActivitySQL{Req: req, Act: act, SQL: sql, Res: r}
	return NewResult(act.Name()+" result", nil, req, act, r, page), nil
}
