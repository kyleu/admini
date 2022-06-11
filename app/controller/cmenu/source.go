package cmenu

import (
	"context"

	"admini.dev/admini/app"
	"admini.dev/admini/app/lib/menu"
	"admini.dev/admini/app/util"
)

func sourceItems(ctx context.Context, as *app.State, logger util.Logger) menu.Items {
	ss, err := as.Services.Sources.List(logger)
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
