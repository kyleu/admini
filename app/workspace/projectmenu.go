package workspace

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"admini.dev/admini/app"
	"admini.dev/admini/app/action"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/lib/schema"
	"admini.dev/admini/app/lib/schema/model"
	"admini.dev/admini/app/project"
	"admini.dev/admini/app/util"
)

func ProjectMenu(as *app.State, prj *project.View) (menu.Items, error) {
	pKey := prj.Project.Key
	rt := "/x"
	if strings.HasPrefix(pKey, project.SourceProjectPrefix) {
		pKey = strings.TrimPrefix(pKey, project.SourceProjectPrefix)
		rt = "/s"
	}

	m, err := ToMenu(as, fmt.Sprintf("%s/%s", rt, pKey), prj.Project.Actions, prj)
	if err != nil {
		return nil, err
	}

	const configDesc = "Edit this project's settings and actions"
	config := &menu.Item{Key: "settings", Description: configDesc, Icon: "cog"}
	if strings.HasPrefix(prj.Project.Key, project.SourceProjectPrefix) {
		config.Route = fmt.Sprintf("/source/%s", strings.TrimPrefix(prj.Project.Key, project.SourceProjectPrefix))
		config.Title = "Source Settings"
	} else {
		config.Route = fmt.Sprintf("/project/%s", prj.Project.Key)
		config.Title = "Project Settings"
	}

	ret := append(menu.Items{}, m...)
	ret = append(ret, menu.Separator, config, menuItemBack)

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
			Icon:        act.IconWithFallback(),
			Route:       p,
		}
		var err error
		switch act.TypeKey {
		case action.TypeFolder.Key:
			// noop
		case action.TypeStatic.Key:
			// noop
		case action.TypeSeparator.Key:
			x = &menu.Item{}

		case action.TypeAll.Key:
			err = itemsForAll(x, view)
		case action.TypeSource.Key:
			err = itemsForSource(x, act, view)
		case action.TypePackage.Key:
			err = itemsForPackage(x, act, view)
		case action.TypeModel.Key:
			// err = itemsForModel(x, act, view)
		case action.TypeActivity.Key:
			// noop

		default:
			err = errors.Errorf("unhandled menu action type [%s]", act.TypeKey)
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

func itemsForAll(x *menu.Item, view *project.View) error {
	for _, src := range view.Sources {
		sch, err := view.Schemata.GetWithError(src.Key)
		if err != nil {
			return err
		}

		path := filepath.Join(x.Route, src.Key)
		kid := &menu.Item{
			Key:         src.Key,
			Title:       src.Name(),
			Icon:        src.IconWithFallback(),
			Description: src.Description,
			Route:       path,
			Children:    SourceMenuPackage(sch.ModelsByPackage(), path),
		}
		x.Children = append(x.Children, kid)
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
	pkgStr, err := act.Config.GetString(action.TypePackage.Key, false)
	if err != nil {
		return errors.Wrap(err, "no [package] in config")
	}
	pkg := util.StringSplitAndTrim(pkgStr, "/")
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

func schemaFor(act *action.Action, view *project.View) (*schema.Schema, error) {
	sourceKey, err := act.Config.GetString(action.TypeSource.Key, false)
	if err != nil {
		return nil, errors.Wrap(err, "key [source] was not provided")
	}

	sch, err := view.Schemata.GetWithError(sourceKey)
	if err != nil {
		return nil, errors.Wrapf(err, "source [%s] is not included in this project", sourceKey)
	}
	return sch, nil
}
