package action

import (
	"github.com/kyleu/admini/app/schema"
	model2 "github.com/kyleu/admini/app/schema/model"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
)

func GetSource(act *Action, schemata schema.Schemata) (*model2.Package, error) {
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

func GetItem(act *Action, schemata schema.Schemata, key string) (interface{}, []string, error) {
	sch, err := GetSource(act, schemata)
	if err != nil {
		return nil, nil, err
	}
	p, err := act.Config.GetString(key, false)
	if err != nil {
		return nil, nil, errors.Wrap(err, "config key [package] must be provided")
	}

	i, remaining := sch.Get(util.SplitAndTrim(p, "/"))
	if i == nil {
		return nil, nil, errors.Errorf("no item found at path [%s]", p)
	}
	return i, remaining, nil
}

func GetPackage(act *Action, schemata schema.Schemata) (*model2.Package, []string, error) {
	i, remaining, err := GetItem(act, schemata, TypePackage.Key)
	if err != nil {
		return nil, nil, err
	}
	switch t := i.(type) {
	case *model2.Package:
		return t, remaining, nil
	default:
		return nil, nil, errors.Errorf("unhandled type for package: %T", t)
	}
}

func GetModel(act *Action, schemata schema.Schemata) (*model2.Model, []string, error) {
	i, remaining, err := GetItem(act, schemata, TypeModel.Key)
	if err != nil {
		return nil, nil, err
	}
	switch t := i.(type) {
	case *model2.Model:
		return t, remaining, nil
	default:
		return nil, nil, errors.Errorf("unhandled type for model: %T", t)
	}
}
