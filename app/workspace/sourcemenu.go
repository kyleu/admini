package workspace

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/schema"
	"github.com/kyleu/admini/app/util"
)

var menuItemBack = &menu.Item{
	Key:         "_back",
	Title:       "Back to " + util.AppName,
	Description: "Return to " + util.AppName,
	Route:       "/",
}

func SourceMenu(as *app.State, source string, sch *schema.Schema) menu.Items {
	ret := menu.Items{
		{
			Key:         "overview",
			Title:       "Project overview",
			Description: "Overview of the data source, displaying details about the configuration",
			Route:       as.Route("workspace.source", "key", source),
		},
		menu.Separator,
	}

	mp := sch.ModelsByPackage()

	for _, m := range mp.ChildModels {
		ret = sourceMenuAddModel(as, source, ret, m)
	}

	for _, p := range mp.ChildPackages {
		ret = sourceMenuAddPackage(as, source, ret, p, []string{})
	}

	ret = append(ret, menu.Separator, menuItemBack)

	return ret
}

func sourceMenuAddModel(as *app.State, source string, ret menu.Items, m *model.Model) menu.Items {
	return append(ret, &menu.Item{
		Key:         m.Key,
		Title:       m.Key,
		Description: m.Type.String() + " model [" + m.Key + "]",
		Route:       as.Route("workspace.source", "key", source) + "/" + m.PathString(),
	})
}

func sourceMenuAddPackage(as *app.State, source string, ret menu.Items, mp *model.Package, path []string) menu.Items {
	path = append(path, mp.Key)
	desc := fmt.Sprintf("package [%v], containing [%v] models", mp.Key, len(mp.ChildModels))

	if len(mp.ChildPackages) > 0 {
		desc += fmt.Sprintf(" and [%v] packages", len(mp.ChildPackages))
	}
	i := &menu.Item{
		Key:         mp.Key,
		Title:       mp.Key,
		Description: desc,
		Route:       as.Route("workspace.source", "key", source) + "/" + strings.Join(path, "/"),
	}

	for _, m := range mp.ChildModels {
		i.Children = sourceMenuAddModel(as, source, i.Children, m)
	}

	for _, p := range mp.ChildPackages {
		i.Children = sourceMenuAddPackage(as, source, i.Children, p, path)
	}

	return append(ret, i)
}
