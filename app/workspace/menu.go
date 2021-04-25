package workspace

import (
	"fmt"
	"strings"

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

func SchemaMenu(as *app.State, source string, sch *schema.Schema) menu.Items {
	ret := menu.Items{
		{
			Key:         "overview",
			Title:       "Project overview",
			Description: "Overview of the project, displaying details about the configuration",
			Route:       as.Route("workspace", "key", source),
		},
	}

	mp := sch.ModelsByPackage()

	for _, m := range mp.ChildModels {
		ret = menuAddModel(as, source, ret, m)
	}

	for _, p := range mp.ChildPackages {
		ret = menuAddPackage(as, source, ret, p, []string{})
	}

	ret = append(ret, menuItemBack)

	return ret
}

func menuAddModel(as *app.State, source string, ret menu.Items, m *schema.Model) menu.Items {
	return append(ret, &menu.Item{
		Key:         m.Key,
		Title:       m.Key,
		Description: m.Type.String() + " model [" + m.Key + "]",
		Route:       as.Route("workspace", "key", source) + m.Path(),
	})
}

func menuAddPackage(as *app.State, source string, ret menu.Items, mp *schema.ModelPackage, path []string) menu.Items {
	path = append(path, mp.Key)
	desc := fmt.Sprintf("package [%v], containing [%v] models", mp.Key, len(mp.ChildModels))

	if len(mp.ChildPackages) > 0 {
		desc += fmt.Sprintf(" and [%v] packages", len(mp.ChildPackages))
	}
	i := &menu.Item{
		Key:         mp.Key,
		Title:       mp.Key,
		Description: desc,
		Route:       as.Route("workspace", "key", source) + "/" + strings.Join(path, "/"),
	}

	for _, m := range mp.ChildModels {
		i.Children = menuAddModel(as, source, i.Children, m)
	}

	for _, p := range mp.ChildPackages {
		i.Children = menuAddPackage(as, source, i.Children, p, path)
	}

	return append(ret, i)
}
