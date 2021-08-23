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

func MenuFor(ctx context.Context, as *app.State) (menu.Items, error) {
	return menu.Items{
		&menu.Item{Key: "projects", Title: "Projects", Description: "Projects!", Icon: "star", Route: "/project", Children: projectItems(as)},
		menu.Separator,
		&menu.Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Icon: "database", Route: "/source", Children: sourceItems(as)},
		menu.Separator,
		&menu.Item{Key: "sandbox", Title: "Sandboxes", Description: "Playgrounds for testing new features", Icon: "social", Route: "/sandbox", Children: sandboxItems()},
		menu.Separator,
		&menu.Item{Key: "settings", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/settings"},
		&menu.Item{Key: "refresh", Title: "Refresh", Description: "Reload all cached in " + util.AppName, Icon: "refresh", Route: "/refresh"},
		&menu.Item{Key: "about", Title: "About", Description: "Get assistance and advice for using " + util.AppName, Icon: "question", Route: "/about"},
	}, nil
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

func sandboxItems() menu.Items {
	ret := make(menu.Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &menu.Item{
			Key:         s.Key,
			Title:       s.Title,
			Icon:        s.Icon,
			Description: fmt.Sprintf("Sandbox [%s]", s.Key),
			Route:       fmt.Sprintf("/sandbox/%s", s.Key),
		})
	}
	return ret
}
