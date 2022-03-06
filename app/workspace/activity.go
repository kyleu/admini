package workspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"admini.dev/app"
	"admini.dev/app/action"
	"admini.dev/app/controller/cutil"
	"admini.dev/views/vaction"
)

func sourceActivity(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
	a := act.Config.GetStringOpt(action.TypeActivity.Key)
	switch a {
	case "sql":
		return sourceActivitySQL(req, act, as)
	case "":
		return ErrResult(req, act, errors.New("must provide activity in action config"))
	default:
		return ErrResult(req, act, errors.Errorf("invalid activity [%s] in action config", a))
	}
}

func sourceActivitySQL(req *cutil.WorkspaceRequest, act *action.Action, as *app.State) (*Result, error) {
	srcKey := act.Config.GetStringOpt(action.TypeSource.Key)
	if srcKey == "" {
		return ErrResult(req, act, errors.New("must provide source in action config"))
	}
	sql := act.Config.GetStringOpt("query")
	if string(req.RC.Method()) == fasthttp.MethodPost {
		frm, err := cutil.ParseForm(req.RC)
		if err != nil {
			return nil, errors.Wrap(err, "unable to parse form")
		}

		s := frm.GetStringOpt("sql")
		if s != "" {
			sql = s
		}
	}

	if sql == "" {
		page := &vaction.ActivitySQL{Req: req, Act: act, SQL: sql, Res: nil}
		return NewResult("", nil, req, act, act, page), nil
	}
	src := req.Sources.Get(srcKey)
	if src == nil {
		return ErrResult(req, act, errors.Errorf("source [%s] is not in this project's configuration", srcKey))
	}

	ld, err := as.Services.Loaders.Get(src.Type, src.Key, src.Config)
	if err != nil {
		return ErrResult(req, act, errors.New("unable to create loader"))
	}

	r, err := ld.Query(req.Context, nil, sql)
	if err != nil {
		return ErrResult(req, act, errors.New("unable to execute query"))
	}
	page := &vaction.ActivitySQL{Req: req, Act: act, SQL: sql, Res: r}
	return NewResult(act.Name()+" result", nil, req, act, r, page), nil
}
