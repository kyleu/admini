package menu

import (
	"github.com/kyleu/admini/app/sandbox"
	"github.com/kyleu/admini/app/source"
	"github.com/kyleu/admini/app/util"
)

var (
	itemSeparator = &Item{}
	itemSandbox   = &Item{Key: "sandbox", Title: "Sandboxes", Description: "Playgrounds for testing new features", Route: "/sandbox", Children: sandboxItems()}
	itemTest      = &Item{Key: "test", Title: "Tests", Description: "Tests!", Route: "/test", Children: testItems()}
	itemSettings  = &Item{Key: "settings", Title: "Settings", Description: "System-wide settings and preferences", Route: "/settings"}
	itemModules   = &Item{Key: "modules", Title: "Modules", Description: "Lists the Go modules used by " + util.AppName, Route: "/modules"}
	itemRoutes    = &Item{Key: "routes", Title: "Routes", Description: "Lists the available HTTP routes", Route: "/routes"}
	itemFeedback  = &Item{Key: "feedback", Title: "Send feedback", Description: "Submit feedback so we can improve " + util.AppName, Route: "/feedback"}
	itemHelp      = &Item{Key: "help", Title: "Help", Description: "Get assistance and advice for using " + util.AppName, Route: "/help"}
)

func For(sources *source.Service) Items {
	var itemSources = &Item{Key: "sources", Title: "Sources", Description: "Sources of data, used as input", Route: "/source", Children: sourceItems(sources)}
	return Items{
		itemSandbox, itemSeparator, itemSources, itemSeparator, itemTest,
		itemSeparator, itemSettings, itemModules, itemRoutes, itemFeedback, itemHelp,
	}
}

func sandboxItems() Items {
	ret := make(Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Title,
			Description: "Sandbox [" + s.Key + "]",
			Route:       "/sandbox/" + s.Key,
		})
	}
	return ret
}

func sourceItems(sources *source.Service) Items {
	ss, err := sources.List()
	if err != nil {
		return Items{{Key: "error", Title: "Error", Description: err.Error()}}
	}

	ret := make(Items, 0, len(ss))
	for _, s := range ss {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Title,
			Description: s.Description,
			Route:       "/source/" + s.Key,
		})
	}

	return ret
}

func testItems() Items {
	return Items{
		{
			Key:   "a",
			Title: "Test A",
			Route: "/test/a",
			Children: Items{
				{
					Key:   "a1",
					Title: "Test A1",
					Route: "/test/a/1",
					Children: Items{
						{
							Key:   "a1x",
							Title: "Test A1X",
							Route: "/test/a/1/x",
						},
					},
				},
			},
		},
		{
			Key:   "b",
			Title: "Test B",
			Route: "/test/b",
			Children: Items{
				{
					Key:   "b1",
					Title: "Test B1",
					Route: "/test/b/1",
					Children: Items{
						{
							Key:   "b1x",
							Title: "Test B1X",
							Route: "/test/b/1/x",
						},
					},
				},
			},
		},
	}
}
