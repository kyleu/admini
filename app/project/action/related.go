package action

import (
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func getSource(act *Action, schemata schema.Schemata) (*model.Package, error) {
	k, err := act.Config.GetString("source", false)
	if err != nil {
		return nil, errors.Wrap(err, "config key [source] must be provided")
	}
	sch, err := schemata.GetWithError(k)
	if err != nil {
		return nil, errors.Wrapf(err, "can't find source in project with key [%s]", k)
	}
	return sch.ModelsByPackage(), nil
}

func getPackage(act *Action, schemata schema.Schemata) (*model.Package, []string, error) {
	sch, err := getSource(act, schemata)
	if err != nil {
		return nil, nil, err
	}
	p, err := act.Config.GetString(TypePackage.Key, false)
	if err != nil {
		return nil, nil, errors.Wrap(err, "config key [package] must be provided")
	}

	i, remaining := sch.Get(util.SplitAndTrim(p, "/"))
	if i == nil {
		return nil, nil, errors.Errorf("no package found at path [%s]", p)
	}

	switch t := i.(type) {
	case *model.Package:
		return t, remaining, nil
	default:
		return nil, nil, errors.Errorf("unhandled type for package: %T", t)
	}
}

func getModel(act *Action, schemata schema.Schemata) (*model.Model, []string, error) {
	sch, err := getSource(act, schemata)
	if err != nil {
		return nil, nil, err
	}
	p, err := act.Config.GetString(TypeModel.Key, false)
	if err != nil {
		return nil, nil, errors.Wrap(err, "config key [model] must be provided")
	}

	i, remaining := sch.Get(util.SplitAndTrim(p, "/"))
	if i == nil {
		return nil, nil, errors.Errorf("no model found at path [%s]", p)
	}

	switch t := i.(type) {
	case *model.Model:
		return t, remaining, nil
	default:
		return nil, nil, errors.Errorf("unhandled type for model: %T", t)
	}
}
