package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
	"github.com/pkg/errors"
	"path/filepath"
)

func ProjectMenu(as *app.State, prj *project.View) (menu.Items, error) {
	ret := menu.Items{
		{
			Key:         "overview",
			Title:       "Project overview",
			Description: "Overview of the project, displaying details about the configuration",
			Route:       as.Route("workspace", "key", prj.Project.Key),
		},
		menu.Separator,
	}

	m, err := ToMenu(as, as.Route("workspace", "key", prj.Project.Key), prj.Project.Actions, prj)
	if err != nil {
		return nil, err
	}

	ret = append(ret, m...)
	ret = append(ret, menu.Separator, menuItemBack)

	return ret, nil
}

func ToMenu(as *app.State, path string, a action.Actions, view *project.View) (menu.Items, error) {
	ret := make(menu.Items, 0, len(a))
	for _, act := range a {
		p := filepath.Join(path, act.Key)
		x := &menu.Item{
			Key:         act.Key,
			Title:       act.Name(),
			Description: act.Description,
			Icon:        act.Icon,
			Route:       p,
		}
		var err error
		switch act.Type {
		case "":
			// noop
		case action.TypeAll.Key:
			err = itemsForAll(x, act, view)
		case action.TypeSource.Key:
			err = itemsForSource(x, act, view)
		case action.TypePackage.Key:
			err = itemsForPackage(x, act, view)
		case action.TypeModel.Key:
			err = itemsForModel(x, act, view)
		case action.TypeStatic.Key:
			// noop
		default:
			err = errors.New("unhandled action type [" + act.Type + "]")
		}
		if err != nil {
			return nil, err
		}

		if len(act.Children) > 0 {
			kids, err := ToMenu(as, p, act.Children, view)
			if err != nil {
				return nil, err
			}
			x.Children = append(x.Children, kids...)
		}
		ret = append(ret, x)
	}

	return ret, nil
}

func itemsForAll(x *menu.Item, act *action.Action, view *project.View) error {
	for _, src := range view.Sources {
		x.Children = append(x.Children, &menu.Item{
			Key:         src.Key,
			Title:       src.Name(),
			Description: src.Description,
			Route:       "TODO",
		})
	}
	return nil
}

func itemsForSource(x *menu.Item, act *action.Action, view *project.View) error {
	sch, err := schemaFor(act, view)
	if err != nil {
		return err
	}
	x.Children = SourceMenuPackage(sch.ModelsByPackage(), x.Route)
	return nil
}

func itemsForPackage(x *menu.Item, act *action.Action, view *project.View) error {
	sch, err := schemaFor(act, view)
	if err != nil {
		return err
	}
	pkgStr, ok := act.Config["package"]
	if !ok {
		return errors.New("no [package] in config")
	}
	pkg := util.SplitAndTrim(pkgStr, "/")
	if len(pkg) == 0 {
		return errors.New("config [package] is empty")
	}
	i, _ := sch.ModelsByPackage().Get(pkg)
	switch t := i.(type) {
	case nil:
		return errors.New("config [package] must refer to an existing package")
	case *model.Package:
		x.Children = SourceMenuPackage(t, x.Route)
		return nil
	default:
		return errors.New("config [package] must refer to a package")
	}
}

func itemsForModel(x *menu.Item, act *action.Action, view *project.View) error {
	sch, err := schemaFor(act, view)
	if err != nil {
		return err
	}
	pkgStr, ok := act.Config["model"]
	if !ok {
		return errors.New("no [model] in config")
	}
	mPath := util.SplitAndTrim(pkgStr, "/")
	if len(mPath) == 0 {
		return errors.New("config [model] is empty")
	}
	i, _ := sch.ModelsByPackage().Get(mPath)
	switch t := i.(type) {
	case nil:
		return errors.New("config [model] must refer to an existing model")
	case *model.Model:
		x.Children = menu.Items{
			{Key: "test", Title: "Test " + t.Name(), Route: act.Pkg.ToPath(act.Key)},
		}
		return nil
	default:
		return errors.New("config [package] must refer to a package")
	}
}

func schemaFor(act *action.Action, view *project.View) (*schema.Schema, error) {
	sourceKey, ok := act.Config["source"]
	if !ok {
		return nil, errors.New("source [" + sourceKey + "] is not included in this project")
	}

	sch, err := view.Schemata.GetWithError(sourceKey)
	if err != nil {
		return nil, errors.Wrap(err, "source ["+sourceKey+"] is not included in this project")
	}
	return sch, nil
}
