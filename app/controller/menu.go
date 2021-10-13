// Package controller $PF_IGNORE$
package controller

import (
	"context"
	"fmt"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/sandbox"
	"github.com/kyleu/admini/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State) (menu.Items, error) {
	ret := menu.Items{
		&menu.Item{Key: "projects", Title: "Projects", Description: "Projects!", Icon: "star", Route: "/project", Children: projectItems(as)},
		menu.Separator,
		&menu.Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Icon: "database", Route: "/source", Children: sourceItems(as)},
		menu.Separator,
	}
	if isAdmin {
		ret = append(ret,
			sandbox.Menu(),
			menu.Separator,
			&menu.Item{Key: "settings", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin/settings"},
			&menu.Item{Key: "refresh", Title: "Refresh", Description: "Reload all cached in " + util.AppName, Icon: "refresh", Route: "/refresh"},
		)
	}
	aboutDesc := "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	return ret, nil
}

func projectItems(as *app.State) menu.Items {
	ps, err := as.Services.Projects.List()
	if err != nil {
		return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	ret := make(menu.Items, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, &menu.Item{
			Key:         p.Key,
			Title:       p.Name(),
			Icon:        p.IconWithFallback(),
			Description: p.Description,
			Route:       fmt.Sprintf("/project/%s", p.Key),
		})
	}
	return ret
}

func sourceItems(as *app.State) menu.Items {
	ss, err := as.Services.Sources.List()
	if err != nil {
		return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	ret := make(menu.Items, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, &menu.Item{
			Key:         s.Key,
			Title:       s.Name(),
			Icon:        s.IconWithFallback(),
			Description: s.Description,
			Route:       fmt.Sprintf("/source/%s", s.Key),
		})
	}
	return ret
}
