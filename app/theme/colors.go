package theme

import (
	"fmt"
	"strings"
)

type Colors struct {
	Foreground      string `json:"fg"`
	ForegroundMuted string `json:"fgm"`
	Background      string `json:"bg"`
	BackgroundMuted string `json:"bgm"`

	Link        string `json:"l"`
	LinkVisited string `json:"lv"`

	NavForeground string `json:"nf"`
	NavBackground string `json:"nb"`

	MenuForeground         string `json:"mf"`
	MenuBackground         string `json:"mb"`
	MenuBackgroundSelected string `json:"mbs"`
}

func (c *Colors) CSS(indent int) string {
	sb := &strings.Builder{}
	addLine(sb, ":root {", indent)
	prop := func(k string, v string) {
		addLine(sb, fmt.Sprintf("--color-%s: %s;", k, v), indent+1)
	}
	prop("foreground", c.Foreground)
	prop("foreground-muted", c.ForegroundMuted)
	prop("background", c.Background)
	prop("background-muted", c.BackgroundMuted)
	addLine(sb, "", 0)
	prop("link", c.Link)
	prop("link-visited", c.LinkVisited)
	addLine(sb, "", 0)
	prop("nav-foreground", c.NavForeground)
	prop("nav-background", c.NavBackground)
	addLine(sb, "", 0)
	prop("menu-foreground", c.MenuForeground)
	prop("menu-background", c.MenuBackground)
	prop("menu-background-selected", c.MenuBackgroundSelected)
	addLine(sb, "}", indent)
	return sb.String()
}
