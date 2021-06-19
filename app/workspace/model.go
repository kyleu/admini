package workspace

import (
	"strings"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/action"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func processModel(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, additional []string) (*Result, error) {
	if len(additional) == 0 {
		return processModelList(req, act, srcKey, m)
	}
	switch additional[0] {
	case "new":
		if string(req.Ctx.Method()) == fasthttp.MethodPost {
			return processModelAdd(req, act, srcKey, m, additional)
		}
		return processModelNew(req, act, srcKey, m)
	case "export":
		return processModelExport(req, act, m, additional[1:])
	case "v":
		return processModelView(req, act, srcKey, m, additional[1:])
	case "x":
		if string(req.Ctx.Method()) == fasthttp.MethodPost {
			return processModelSave(req, act, srcKey, m, additional[1:])
		}
		return processModelEdit(req, act, srcKey, m, additional[1:])
	case "d":
		return processModelDelete(req, act, srcKey, m, additional[1:])
	default:
		return nil, errors.Errorf("unhandled model parameters [%s]", strings.Join(additional, "/"))
	}
}

func getModel(m *model.Model, idStrings []string, ld loader.Loader) ([]interface{}, error) {
	ids := make([]interface{}, 0, len(idStrings))
	for _, x := range idStrings {
		ids = append(ids, x)
	}

	rs, err := ld.Get(m, ids)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve model [%s] with key [%s]", m.Path().String(), strings.Join(idStrings, "/"))
	}
	switch len(rs.Data) {
	case 0:
		return nil, errors.Errorf("no [%s] found with key [%s]", m.Name(), strings.Join(idStrings, "/"))
	case 1:
		return rs.Data[0], nil
	default:
		return nil, errors.Errorf("multiple [%s] found with key [%s]", m.Name(), strings.Join(idStrings, "/"))
	}
}

func loaderFor(req *cutil.WorkspaceRequest, srcKey string) (*source.Source, loader.Loader, util.ParamSet, error) {
	s, err := req.Sources.GetWithError(srcKey)
	if err != nil {
		return nil, nil, nil, err
	}

	l, err := req.AS.Loaders.Get(s.Type, s.Key, s.Config)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "no loader available")
	}
	p := cutil.ParamSetFromRequest(req.Ctx)
	return s, l, p, nil
}
