// Package theme - Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"admini.dev/admini/app/util"
)

var Default = func() *Theme {
	nbl := "#8eb6bd"
	if o := util.GetEnv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#495f64"
	if o := util.GetEnv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key: "default",
		Light: &Colors{
			Border: "1px solid #dddddd", LinkDecoration: "none",
			Foreground: "#000000", ForegroundMuted: "#999999",
			Background: "#ffffff", BackgroundMuted: "#eeeeee",
			LinkForeground: "#2d414e", LinkVisitedForeground: "#406379",
			NavForeground: "#000000", NavBackground: nbl,
			MenuForeground: "#000000", MenuSelectedForeground: "#000000",
			MenuBackground: "#c1e3e8", MenuSelectedBackground: "#8eb6bd",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#ffffff", ForegroundMuted: "#999999",
			Background: "#121212", BackgroundMuted: "#333333",
			LinkForeground: "#2d414e", LinkVisitedForeground: "#aaaaaa",
			NavForeground: "#ffffff", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuSelectedForeground: "#ffffff",
			MenuBackground: "#171f24", MenuSelectedBackground: "#333333",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()
