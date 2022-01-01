package workspace

import (
	"context"
	"strings"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/lib/filter"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/admini/app/action"

	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/loader"
	"github.com/kyleu/admini/app/source"
	"github.com/pkg/errors"
)

func processModel(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, additional []string, as *app.State) (*Result, error) {
	if len(additional) == 0 {
		return processModelList(req, act, srcKey, m, as)
	}
	switch additional[0] {
	case "new":
		if string(req.Ctx.Method()) == fasthttp.MethodPost {
			return processModelAdd(req, act, srcKey, m, additional[1:], as)
		}
		return processModelNew(req, act, srcKey, m, as)
	case "v":
		return processModelView(req, act, srcKey, m, additional[1:], as)
	case "x":
		if string(req.Ctx.Method()) == fasthttp.MethodPost {
			return processModelSave(req, act, srcKey, m, additional[1:], as)
		}
		return processModelEdit(req, act, srcKey, m, additional[1:], as)
	case "d":
		return processModelDelete(req, act, srcKey, m, additional[1:], as)
	default:
		return nil, errors.Errorf("unhandled model parameters [%s]", strings.Join(additional, "/"))
	}
}

func getModel(ctx context.Context, m *model.Model, idStrings []string, ld loader.Loader) ([]interface{}, error) {
	ids := make([]interface{}, 0, len(idStrings))
	for _, x := range idStrings {
		ids = append(ids, x)
	}

	rs, err := ld.Get(ctx, m, ids)
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

func loaderFor(req *cutil.WorkspaceRequest, srcKey string, as *app.State) (*source.Source, loader.Loader, error) {
	s, err := req.Sources.GetWithError(srcKey)
	if err != nil {
		return nil, nil, err
	}

	l, err := as.Services.Loaders.Get(s.Type, s.Key, s.Config)
	if err != nil {
		return nil, nil, errors.Wrap(err, "no loader available")
	}
	return s, l, nil
}

func optionsFor(req *cutil.WorkspaceRequest) filter.OptionsMap {
	p := cutil.ParamSetFromRequest(req.Ctx)
	ret := filter.OptionsMap{}
	for k, v := range p {
		ret[k] = &filter.Options{Sort: nil, Filter: nil, Group: nil, Params: v}
	}
	return ret
}
