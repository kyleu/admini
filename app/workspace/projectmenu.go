package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/project"
	"github.com/kyleu/admini/app/project/action"
	"path/filepath"
)

func ProjectMenu(as *app.State, prj *project.Project) menu.Items {
	ret := menu.Items{
		{
			Key:         "overview",
			Title:       "Project overview",
			Description: "Overview of the project, displaying details about the configuration",
			Route:       as.Route("workspace", "key", prj.Key),
		},
		menu.Separator,
	}

	ret = append(ret, ToMenu(as.Route("workspace", "key", prj.Key), prj.Actions)...)
	ret = append(ret, menu.Separator, menuItemBack)

	return ret
}

func ToMenu(path string, a action.Actions) (menu.Items) {
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
		if len(act.Children) > 0 {
			x.Children = ToMenu(p, act.Children)
		}
		ret = append(ret, x)
	}

	return ret
}

