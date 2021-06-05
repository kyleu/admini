package user

import (
	"fmt"
	"strings"
)

type Colors struct {
	Foreground      string
	ForegroundMuted string
	Background      string
	BackgroundMuted string

	Highlight   string
	Link        string
	LinkVisited string

	NavForeground string
	NavBackground string

	MenuForeground         string
	MenuBackground         string
	MenuBackgroundSelected string
}

func (c *Colors) CSS(indent int) string {
	sb := &strings.Builder{}
	add(sb, ":root {", indent)
	prop := func(k string, v string) {
		add(sb, fmt.Sprintf("--color-%s: %s;", k, v), indent + 1)
	}
	prop("foreground", c.Foreground)
	prop("foreground-muted", c.ForegroundMuted)
	prop("background", c.Background)
	prop("background-muted", c.BackgroundMuted)
	add(sb, "", 0)
	prop("highlight", c.Highlight)
	prop("link", c.Link)
	prop("link-visited", c.LinkVisited)
	add(sb, "", 0)
	prop("nav-foreground", c.NavForeground)
	prop("nav-background", c.NavBackground)
	add(sb, "", 0)
	prop("menu-foreground", c.MenuForeground)
	prop("menu-background", c.MenuBackground)
	prop("menu-background-selected", c.MenuBackgroundSelected)
	add(sb, "}", indent)
	return sb.String()
}
