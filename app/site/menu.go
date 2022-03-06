// Content managed by Project Forge, see [projectforge.md] for details.
package site

import (
	"context"

	"admini.dev/app"
	"admini.dev/app/lib/menu"
	"admini.dev/app/lib/user"
)

const (
	keyInstall     = "install"
	keyDownload    = "download"
	keyContrib     = "contributing"
	keyCustomizing = "customizing"
	keyTech        = "technology"
)

func Menu(ctx context.Context, as *app.State, _ *user.Profile, _ user.Accounts) menu.Items {
	return menu.Items{
		{Key: keyInstall, Title: "Install", Icon: "code", Route: "/" + keyInstall},
		{Key: keyDownload, Title: "Download", Icon: "download", Route: "/" + keyDownload},
		{Key: keyCustomizing, Title: "Customizing", Icon: "code", Route: "/" + keyCustomizing},
		{Key: keyContrib, Title: "Contributing", Icon: "cog", Route: "/" + keyContrib},
		{Key: keyTech, Title: "Technology", Icon: "cog", Route: "/" + keyTech},
	}
}
