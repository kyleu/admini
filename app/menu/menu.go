package menu

import (
	"github.com/kyleu/admini/app/sandbox"
)

var Menu = Items{
	{Key: "sandbox", Title: "Sandboxes", Route: "/sandbox", Children: sandboxItems()},
	{},
	{Key: "settings", Title: "Settings", Route: "/settings"},
	{Key: "feedback", Title: "Send feedback", Route: "/feedback"},
	{Key: "help", Title: "Help", Route: "/help"},
}

func sandboxItems() Items {
	ret := make(Items, 0, len(sandbox.AllSandboxes))
	for _, s := range sandbox.AllSandboxes {
		ret = append(ret, &Item{
			Key:         s.Key,
			Title:       s.Key,
			Route:       "/sandbox/" + s.Key,
		})
	}
	return ret
}
