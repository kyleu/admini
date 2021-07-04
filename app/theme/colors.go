package theme

import (
	"fmt"
	"strings"
)

type Colors struct {
	Border         string `json:"brd"`
	LinkDecoration string `json:"ld"`

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

	ModalBackdrop string `json:"mbd"`
	Success       string `json:"ok"`
	Error         string `json:"err"`
}

func (c *Colors) CSS(key string, indent int) string {
	sb := &strings.Builder{}
	addLine(sb, key+" {", indent)
	prop := func(k string, v string) {
		addLine(sb, fmt.Sprintf("--%s: %s;", k, v), indent+1)
	}
	prop("border", c.Border)
	prop("link-text-decoration", c.LinkDecoration)
	addLine(sb, "", 0)
	prop("color-foreground", c.Foreground)
	prop("color-foreground-muted", c.ForegroundMuted)
	prop("color-background", c.Background)
	prop("color-background-muted", c.BackgroundMuted)
	addLine(sb, "", 0)
	prop("color-link", c.Link)
	prop("color-link-visited", c.LinkVisited)
	addLine(sb, "", 0)
	prop("color-nav-foreground", c.NavForeground)
	prop("color-nav-background", c.NavBackground)
	addLine(sb, "", 0)
	prop("color-menu-foreground", c.MenuForeground)
	prop("color-menu-background", c.MenuBackground)
	prop("color-menu-background-selected", c.MenuBackgroundSelected)
	addLine(sb, "", 0)
	prop("modal-backdrop", c.ModalBackdrop)
	prop("color-success", c.Success)
	prop("color-error", c.Error)
	addLine(sb, "}", indent)
	return sb.String()
}
