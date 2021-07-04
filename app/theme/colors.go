package theme

import (
	"fmt"
	"strings"

	"github.com/kyleu/admini/app/util"
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
	sb.WriteString(key + " {")
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
	return ret.String()
}

func (c *Colors) Clone() *Colors {
	return &Colors{
		Border:                 c.Border,
		LinkDecoration:         c.LinkDecoration,
		Foreground:             c.Foreground,
		ForegroundMuted:        c.ForegroundMuted,
		Background:             c.Background,
		BackgroundMuted:        c.BackgroundMuted,
		Link:                   c.Link,
		LinkVisited:            c.LinkVisited,
		NavForeground:          c.NavForeground,
		NavBackground:          c.NavBackground,
		MenuForeground:         c.MenuForeground,
		MenuBackground:         c.MenuBackground,
		MenuBackgroundSelected: c.MenuBackgroundSelected,
		ModalBackdrop:          c.ModalBackdrop,
		Success:                c.Success,
		Error:                  c.Error,
	}
}

func (c *Colors) ApplyMap(m util.ValueMap, prefix string) *Colors {
	get := func(k string, def string) string {
		x, err := m.GetString(prefix+k, true)
		if err != nil {
			return def
		}
		return x
	}
	c.Border = get("border", c.Border)
	c.LinkDecoration = get("link-decoration", c.LinkDecoration)
	c.Foreground = get("foreground", c.Foreground)
	c.ForegroundMuted = get("foreground-muted", c.ForegroundMuted)
	c.Background = get("background", c.Background)
	c.BackgroundMuted = get("background-muted", c.BackgroundMuted)
	c.Link = get("link", c.Link)
	c.LinkVisited = get("link-visited", c.LinkVisited)
	c.NavForeground = get("nav-foreground", c.NavForeground)
	c.NavBackground = get("nav-background", c.NavBackground)
	c.MenuForeground = get("menu-foreground", c.MenuForeground)
	c.MenuBackground = get("menu-background", c.MenuBackground)
	c.MenuBackgroundSelected = get("menu-background-selected", c.MenuBackgroundSelected)
	c.ModalBackdrop = get("modal-backdrop", c.ModalBackdrop)
	c.Success = get("success", c.Success)
	c.Error = get("error", c.Error)
	return c
}
