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

	css   string
}

func (c *Colors) CSS(key string, indent int) string {
	if c.css != "" {
		return c.css
	}
	sb := &strings.Builder{}
	sb.WriteString(key+" {")
	prop := func(k string, v string) {
		sb.WriteString(fmt.Sprintf(" --%s: %s;", k, v))
	}
	prop("border", c.Border)
	prop("link-text-decoration", c.LinkDecoration)
	prop("color-foreground", c.Foreground)
	prop("color-foreground-muted", c.ForegroundMuted)
	prop("color-background", c.Background)
	prop("color-background-muted", c.BackgroundMuted)
	prop("color-link", c.Link)
	prop("color-link-visited", c.LinkVisited)
	prop("color-nav-foreground", c.NavForeground)
	prop("color-nav-background", c.NavBackground)
	prop("color-menu-foreground", c.MenuForeground)
	prop("color-menu-background", c.MenuBackground)
	prop("color-menu-background-selected", c.MenuBackgroundSelected)
	prop("color-modal-backdrop", c.ModalBackdrop)
	prop("color-success", c.Success)
	prop("color-error", c.Error)
	sb.WriteString("}")
	ret := &strings.Builder{}
	addLine(ret, sb.String(), indent)
	c.css = ret.String()
	return c.css
}
