package menu

import (
	"github.com/kyleu/admini/app/sandbox"
)

var Menu = Items{
	{Key: "sandbox", Title: "Sandboxes", Route: "/sandbox", Children: sandboxItems()},
	{},
	{Key: "test", Title: "Tests", Route: "/test", Children: testItems()},
	{},
	{Key: "settings", Title: "Settings", Route: "/settings"},
	{Key: "modules", Title: "Modules", Route: "/modules"},
	{Key: "routes", Title: "Routes", Route: "/routes"},
	{Key: "feedback", Title: "Send feedback", Route: "/feedback"},
	{Key: "help", Title: "Help", Route: "/help"},
}

func sandboxItems() Items {
	ret := make(Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &Item{
			Key:   s.Key,
			Title: s.Title,
			Route: "/sandbox/" + s.Key,
		})
	}
	return ret
}

func testItems() Items {
	return Items{
		{
			Key:   "testa",
			Title: "Test A",
			Route: "/test/a",
			Children: Items{
				{
					Key:   "testa1",
					Title: "Test A1",
					Route: "/test/a/1",
					Children: Items{
						{
							Key:   "testa1x",
							Title: "Test A1X",
							Route: "/test/a/1/x",
						},
					},
				},
			},
		},
		{
			Key:   "testb",
			Title: "Test B",
			Route: "/test/b",
			Children: Items{
				{
					Key:   "testb1",
					Title: "Test B1",
					Route: "/test/b/1",
					Children: Items{
						{
							Key:   "testb1x",
							Title: "Test B1X",
							Route: "/test/b/1/x",
						},
					},
				},
			},
		},
	}
}
