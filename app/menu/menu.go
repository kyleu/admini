package menu

import (
	"fmt"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/sandbox"
	"github.com/kyleu/admini/app/util"
)

func For(as *app.State) Items {
	return Items{
		&Item{Key: "projects", Title: "Projects", Description: "Projects!", Icon: "star", Route: "/project", Children: projectItems(as)},
		Separator,
		&Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Icon: "database", Route: "/source", Children: sourceItems(as)},
		Separator,
		&Item{Key: "sandbox", Title: "Sandboxes", Description: "Playgrounds for testing new features", Icon: "social", Route: "/sandbox", Children: sandboxItems()},
		Separator,
		&Item{Key: "settings", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/settings"},
		&Item{Key: "refresh", Title: "Refresh", Description: "Reload all cached in " + util.AppName, Icon: "refresh", Route: "/refresh"},
		&Item{Key: "feedback", Title: "Feedback", Description: "Submit feedback so we can improve " + util.AppName, Icon: "mail", Route: "/feedback"},
		&Item{Key: "help", Title: "Help", Description: "Get assistance and advice for using " + util.AppName, Icon: "comment", Route: "/help"},
	}
}

func projectItems(as *app.State) Items {
	ps, err := as.Projects.List()
	if err != nil {
		return Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	ret := make(Items, 0, len(ps))
	for _, p := range ps {
		ret = append(ret, &Item{
			Key:         p.Key,
			Title:       p.Name(),
			Icon:        p.IconWithFallback(),
			Description: p.Description,
			Route:       fmt.Sprintf("/project/%s", p.Key),
		})
	}
	return ret
}

func sourceItems(as *app.State) Items {
	ss, err := as.Sources.List()
	if err != nil {
		return Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	ret := make(Items, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Name(),
			Icon:        s.IconWithFallback(),
			Description: s.Description,
			Route:       fmt.Sprintf("/source/%s", s.Key),
		})
	}
	return ret
}

func sandboxItems() Items {
	ret := make(Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Title,
			Icon:        s.Icon,
			Description: fmt.Sprintf("Sandbox [%s]", s.Key),
			Route:       fmt.Sprintf("/sandbox/%s", s.Key),
		})
	}
	return ret
}
