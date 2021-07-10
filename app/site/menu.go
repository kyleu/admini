package site

import (
	"github.com/kyleu/admini/app/auth"
	"github.com/kyleu/admini/app/menu"
	"github.com/kyleu/admini/app/user"
)

const (
	keyIntro      = "intro"
	keyInstall    = "install"
	keyQuickStart = "quickstart"
	keyContrib    = "contrib"
)

func SiteMenu(p *user.Profile, s auth.Sessions) menu.Items {
	return menu.Items{
		{Key: keyIntro, Title: "Introduction", Icon: "heart", Route: "/intro"},
		{Key: keyInstall, Title: "Install", Icon: "download", Route: "/install"},
		{Key: keyQuickStart, Title: "Quick Start", Icon: "bolt", Route: "/quickstart"},
		{Key: keyContrib, Title: "Contributing", Icon: "cog", Route: "/contrib"},
	}
}
