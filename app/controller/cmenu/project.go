package cmenu

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/util"
)

func projectItems(ctx context.Context, as *app.State, logger util.Logger) menu.Items {
	ps, err := as.Services.Projects.List(ctx, logger)
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
