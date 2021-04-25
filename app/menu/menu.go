package menu

import (
	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/sandbox"
	"github.com/kyleu/admini/app/util"
)

func For(as *app.State) Items {
	return Items{
		&Item{Key: "sandbox", Title: "Sandboxes", Description: "Playgrounds for testing new features", Route: as.Route("sandbox.list"), Children: sandboxItems(as)},
		&Item{},
		&Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Route: as.Route("source.list"), Children: sourceItems(as)},
		&Item{},
		&Item{Key: "settings", Title: "Settings", Description: "System-wide settings and preferences", Route: as.Route("settings")},
		&Item{Key: "modules", Title: "Modules", Description: "Lists the Go modules used by " + util.AppName, Route: as.Route("modules")},
		&Item{Key: "routes", Title: "Routes", Description: "Lists the available HTTP routes", Route: as.Route("routes")},
		&Item{Key: "feedback", Title: "Send feedback", Description: "Submit feedback so we can improve " + util.AppName, Route: as.Route("feedback")},
		&Item{Key: "help", Title: "Help", Description: "Get assistance and advice for using " + util.AppName, Route: as.Route("help")},
	}
}

func sandboxItems(as *app.State) Items {
	ret := make(Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Title,
			Description: "Sandbox [" + s.Key + "]",
			Route:       as.Route("sandbox.run", "key", s.Key),
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
			Title:       s.Title,
			Description: s.Description,
			Route:       as.Route("source.detail", "key", s.Key),
		})
	}

	return ret
}
