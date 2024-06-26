package cmenu

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/filter"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/lib/sandbox"
	"admini.dev/admini/app/lib/telemetry"
	"admini.dev/admini/app/lib/user"
	"admini.dev/admini/app/util"
)

func MenuFor(
	ctx context.Context, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, as *app.State, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	prj := &menu.Item{Key: "projects", Title: "Projects", Description: "Projects!", Icon: "star", Route: "/project", Children: projectItems(ctx, as, logger)}
	srcDesc := "Sources of data"
	src := &menu.Item{Key: "sources", Title: "Sources", Description: srcDesc, Icon: "database", Route: "/source", Children: sourceItems(ctx, as, logger)}
	ret = append(ret, prj, menu.Separator, src, menu.Separator)
	if isAdmin {
		ret = append(ret,
			sandbox.Menu(ctx),
			menu.Separator,
			&menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"},
			&menu.Item{Key: "refresh", Title: "Refresh", Description: "Reload all cached in " + util.AppName, Icon: "refresh", Route: "/refresh"},
		)
	}
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
