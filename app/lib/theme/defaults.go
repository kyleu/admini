package theme

import (
	"fmt"
	"image/color"

	"github.com/pkg/errors"

	"admini.dev/admini/app/util"
)

const (
	white, black = "#ffffff", "#000000"
	threshold    = (65535 * 3) / 2
)

var Default = func() *Theme {
	nbl := "#93aeb3"
	if o := util.GetEnv("app_nav_color_light"); o != "" {
		nbl = o
	}
	nbd := "#102a2e"
	if o := util.GetEnv("app_nav_color_dark"); o != "" {
		nbd = o
	}

	return &Theme{
		Key: "default",
		Light: &Colors{
			Border: "1px solid #dddddd", LinkDecoration: "none",
			Foreground: "#000000", ForegroundMuted: "#2f3d3f",
			Background: "#ffffff", BackgroundMuted: "#e9eeef",
			LinkForeground: "#102326", LinkVisitedForeground: "#102326",
			NavForeground: "#2a2a2a", NavBackground: nbl,
			MenuForeground: "#000000", MenuSelectedForeground: "#000000",
			MenuBackground: "#d3dee0", MenuSelectedBackground: "#a8bec2",
			ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#ff0000",
		},
		Dark: &Colors{
			Border: "1px solid #666666", LinkDecoration: "none",
			Foreground: "#dddddd", ForegroundMuted: "#94a4a7",
			Background: "#121212", BackgroundMuted: "#0f1c1e",
			LinkForeground: "#d3dee0", LinkVisitedForeground: "#93aeb3",
			NavForeground: "#f8f9fa", NavBackground: nbd,
			MenuForeground: "#eeeeee", MenuSelectedForeground: "#dddddd",
			MenuBackground: "#102326", MenuSelectedBackground: "#526c71",
			ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#ff0000",
		},
	}
}()

func TextColorFor(clr string) string {
	c, err := ParseHexColor(clr)
	if err != nil {
		return white
	}
	r, g, b, _ := c.RGBA()
	total := r + g + b
	if total < threshold {
		return white
	}
	return black
}

func ParseHexColor(s string) (color.RGBA, error) {
	ret := color.RGBA{A: 0xff}
	var err error
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &ret.R, &ret.G, &ret.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &ret.R, &ret.G, &ret.B)
		// Double the hex digits:
		ret.R *= 17
		ret.G *= 17
		ret.B *= 17
	default:
		err = errors.Errorf("invalid length [%d], must be 7 or 4", len(s))
	}
	return ret, err
}
