package workspace

import (
	"strings"

	"admini.dev/app"
	"admini.dev/app/action"
	"admini.dev/app/controller/cutil"
	"admini.dev/app/lib/schema/model"
	"admini.dev/app/util"
	"admini.dev/views/vmodel"
	"github.com/pkg/errors"
)

func processModelView(req *cutil.WorkspaceRequest, act *action.Action, srcKey string, m *model.Model, idStrings []string, as *app.State) (*Result, error) {
	_, ld, err := loaderFor(req, srcKey, as)
	if err != nil {
		return ErrResult(req, act, err)
	}

	data, err := getModel(req.Context, m, idStrings, ld)
	if err != nil {
		return ErrResult(req, act, err)
	}

	obj, err := m.OrderedMap(data)
	if err != nil {
		return ErrResult(req, act, err)
	}

	//kids, err := getRelationData(req, m, data, srcKey, as)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to load relation objects")
	//}
	//if len(kids) > 0 {
	//	obj.Append("relatedModels", kids)
	//}

	page := &vmodel.View{Req: req, Act: act, Model: m, Result: data}
	idx := len(req.Path) - len(idStrings) - 1
	if idx < 0 {
		idx = 0
	}

	bc := append(append(act.Path(), req.Path[:idx]...), idStrings...)

	ret := NewResult("", bc, req, act, obj, page)

	return ret, nil
}

func getRelationData(req *cutil.WorkspaceRequest, m *model.Model, data []interface{}, sourceKey string, as *app.State) ([]*util.OrderedMap, error) {
	ret := make([]*util.OrderedMap, 0, len(m.Relationships))
	for _, rel := range m.Relationships {
		rowFK, err := model.GetStrings(m.Fields, rel.SourceFields, data)
		if err != nil {
			return nil, err
		}

		src, relLoader, err := loaderFor(req, sourceKey, as)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to create loader for source [%s]", sourceKey)
		}

		sch, err := req.Schemata.GetWithError(src.Key)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to load schema for source [%s]", src.Key)
		}

		mdl := sch.Models.Get(rel.TargetPkg, rel.TargetModel)
		if mdl == nil {
			const msg = "unable to load model [%s] in package [%s] from schema [%s]"
			return nil, errors.Wrapf(err, msg, rel.TargetModel, rel.TargetPkg.String(), src.Key)
		}

		childData, err := getModel(req.Context, mdl, rowFK, relLoader)
		if err != nil {
			const msg = "unable to load [%s] model with keys [%s] from source [%s]"
			return nil, errors.Wrapf(err, msg, mdl.Key, strings.Join(rel.TargetFields, ", "), src.Key)
		}
		childMap, err := mdl.OrderedMap(childData)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to build [%s] result for relation [%s]", mdl.Key, rel.Key)
		}

		ret = append(ret, childMap)
	}
	return ret, nil
}
