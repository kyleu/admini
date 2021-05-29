package workspace

import (
	"fmt"
	"path/filepath"

	"github.com/kyleu/admini/app/model"

	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/util"
)

var menuItemBack = &menu.Item{
	Key:         "_back",
	Title:       "Back to " + util.AppName,
	Description: "Return to " + util.AppName,
	Route:       "/",
}

func SourceMenuPackage(mp *model.Package, path string) menu.Items {
	ret := menu.Items{}

	for _, m := range mp.ChildModels {
		ret = sourceMenuAddModel(ret, m, path)
	}
	for _, p := range mp.ChildPackages {
		ret = sourceMenuAddPackage(ret, p, path)
	}

	return ret
}

func SourceMenuModel(m *model.Model, path string) *menu.Item {
	return &menu.Item{
		Key:         m.Key,
		Title:       m.Name(),
		Description: m.Description(),
		Route:       path,
	}
}

func sourceMenuAddModel(ret menu.Items, m *model.Model, path string) menu.Items {
	return append(ret, SourceMenuModel(m, filepath.Join(path, m.Key)))
}

func sourceMenuAddPackage(ret menu.Items, mp *model.Package, path string) menu.Items {
	path = filepath.Join(path, mp.Key)
	desc := fmt.Sprintf("package [%v], containing [%v] models", mp.Key, len(mp.ChildModels))

	if len(mp.ChildPackages) > 0 {
		desc += fmt.Sprintf(" and [%v] packages", len(mp.ChildPackages))
	}
	i := &menu.Item{
		Key:         mp.Key,
		Title:       mp.Key,
		Description: desc,
		Route:       path,
	}

	for _, m := range mp.ChildModels {
		i.Children = sourceMenuAddModel(i.Children, m, path)
	}

	for _, p := range mp.ChildPackages {
		i.Children = sourceMenuAddPackage(i.Children, p, path)
	}

	return append(ret, i)
}
