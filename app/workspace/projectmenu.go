package workspace

import (
	"path/filepath"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"
	"github.com/pkg/errors"
)

func ProjectMenu(as *app.State, prj *project.Project) (menu.Items, error) {
	ret := menu.Items{
		{
			Key:         "overview",
			Title:       "Project overview",
			Description: "Overview of the project, displaying details about the configuration",
			Route:       as.Route("workspace", "key", prj.Key),
		},
		menu.Separator,
	}

	m, err := ToMenu(as, as.Route("workspace", "key", prj.Key), prj.Key, prj.Actions, prj.Sources)
	if err != nil {
		return nil, err
	}

	ret = append(ret, m...)
	ret = append(ret, menu.Separator, menuItemBack)

	return ret, nil
}

func ToMenu(as *app.State, path string, prj string, a action.Actions, sources []string) (menu.Items, error) {
	ret := make(menu.Items, 0, len(a))
	for _, act := range a {
		p := filepath.Join(path, act.Key)
		x := &menu.Item{
			Key:         act.Key,
			Title:       act.TitleString(),
			Description: act.Description,
			Icon:        act.Icon,
			Route:       p,
		}
		switch act.Type {
		case action.ActionTypeSource:
			sourceKey, ok := act.Config["source"]
			if !ok {
				return nil, errors.New("source [" + sourceKey + "] is not included in this project")
			}

			ok = false
			for _, s := range sources {
				if s == sourceKey {
					ok = true
				}
			}
			if !ok {
				return nil, errors.New("source [" + sourceKey + "] is not included in this project")
			}

			view, err := as.Projects.LoadView(prj)
			if err != nil {
				return nil, errors.Wrap(err, "can't load project view")
			}

			sch, ok := view.Schemata[sourceKey]
			if !ok {
				return nil, errors.New("schema for source [" + sourceKey + "] was not found")
			}

			x.Children = sourceMenuDetails(sch, x.Route)
		}

		if len(act.Children) > 0 {
			kids, err := ToMenu(as, p, prj, act.Children, sources)
			if err != nil {
				return nil, err
			}
			x.Children = append(x.Children, kids...)
		}
		ret = append(ret, x)
	}

	return ret, nil
}
