package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/action"
	"github.com/kyleu/admini/app/controller/cutil"
	"github.com/kyleu/admini/app/lib/schema/model"
	"github.com/kyleu/admini/app/lib/schema/types"
	"github.com/kyleu/admini/views/vmodel"
)

func processModelList(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, as *app.State) (*Result, error) {
	switch m.Type {
	case model.TypeStruct:
		_, ld, err := loaderFor(req, srcKey, as)
		if err != nil {
			return ErrResult(req, act, err)
		}

		optsMap := optionsFor(req)
		opts := optsMap.Get(m.Key)
		rs, err := ld.List(req.Context, m, opts)
		if err != nil {
			return ErrResult(req, act, errors.Wrapf(err, "unable to list model [%s]", m.Key))
		}
		page := &vmodel.List{Req: req, Act: act, Model: m, Options: opts, Result: rs}
		return NewResult("", nil, req, act, rs, page), nil
	case model.TypeEnum:
		refs := model.Relationships{}
		for _, sch := range req.Schemata {
			for _, mod := range sch.Models {
				for _, f := range mod.Fields {
					if t, ok := f.Type.T.(*types.Enum); ok {
						if t.Ref == m.Key {
							refs = append(refs, &model.Relationship{
								Key:          mod.Key + "_" + f.Key,
								TargetPkg:    mod.Pkg,
								TargetModel:  mod.Key,
								TargetFields: []string{f.Key},
							})
						}
					}
				}
			}
		}
		page := &vmodel.Enum{Req: req, Act: act, Model: m, Refs: refs}
		return NewResult("", nil, req, act, m, page), nil
	default:
		return nil, errors.Errorf("unhandled model type [%s]", m.Type.String())
	}
}
