package workspace

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/project"
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

	ret = append(ret, menu.Separator, menuItemBack)

	return ret
}
