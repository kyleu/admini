// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"context"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/lib/menu"
	"github.com/kyleu/admini/app/lib/sandbox"
	"github.com/kyleu/admini/app/lib/telemetry"
	"github.com/kyleu/admini/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State) (menu.Items, error) {
	ctx, span := telemetry.StartSpan(ctx, "menu", "menu:generate")
	defer span.Complete()

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	projectItems := func(as *app.State) menu.Items {
		ps, err := as.Services.Projects.List(ctx)
		if err != nil {
			return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
		}

		prjMenu := make(menu.Items, 0, len(ps))
		for _, p := range ps {
			prjMenu = append(prjMenu, &menu.Item{
				Key:         p.Key,
				Title:       p.Name(),
				Icon:        p.IconWithFallback(),
				Description: p.Description,
				Route:       "/project/" + p.Key,
			})
		}
		return prjMenu
	}

	sourceItems := func(as *app.State) menu.Items {
		ss, err := as.Services.Sources.List()
		if err != nil {
			return menu.Items{{Key: "error", Title: "Error", Description: err.Error()}}
		}

		srcMenu := make(menu.Items, 0, len(ss))
		for _, s := range ss {
			srcMenu = append(srcMenu, &menu.Item{
				Key:         s.Key,
				Title:       s.Name(),
				Icon:        s.IconWithFallback(),
				Description: s.Description,
				Route:       "/source/" + s.Key,
			})
		}
		return srcMenu
	}
	ret = append(ret,
		&menu.Item{Key: "projects", Title: "Projects", Description: "Projects!", Icon: "star", Route: "/project", Children: projectItems(as)},
		menu.Separator,
		&menu.Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Icon: "database", Route: "/source", Children: sourceItems(as)},
		menu.Separator,
	)
	// $PF_SECTION_END(routes_start)$
	// $PF_INJECT_START(codegen)$
	// $PF_INJECT_END(codegen)$
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		ret = append(ret,
			sandbox.Menu(ctx),
			menu.Separator,
			&menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"},
			&menu.Item{Key: "refresh", Title: "Refresh", Description: "Reload all cached in " + util.AppName, Icon: "refresh", Route: "/refresh"},
		)
	}
	aboutDesc := "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
