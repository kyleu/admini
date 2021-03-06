// Content managed by Project Forge, see [projectforge.md] for details.
package theme

import (
	"io"
	"strings"

	"golang.org/x/exp/slices"

	"admini.dev/admini/app/util"
)

type Theme struct {
	Key   string  `json:"-"`
	Light *Colors `json:"light"`
	Dark  *Colors `json:"dark"`
	css   string
}

func (t *Theme) CSS(indent int) string {
	if t.css != "" {
		return t.css
	}
	sb := &strings.Builder{}
	sb.WriteString("/* theme: " + t.Key + " */\n")
	sb.WriteString(t.Light.CSS(":root", indent))
	sb.WriteString(t.Light.CSS(".mode-light", indent))
	sb.WriteString(t.Dark.CSS(".mode-dark", indent))
	addLine(sb, "", indent)
	addLine(sb, "@media (prefers-color-scheme: dark) {", indent)
	sb.WriteString(t.Dark.CSS(":root", indent+1))
	sb.WriteString(t.Light.CSS(".mode-light", indent+1))
	sb.WriteString(t.Dark.CSS(".mode-dark", indent+1))
	addLine(sb, "}", indent)
	t.css = sb.String()
	return t.css
}

func (t *Theme) Clone(key string) *Theme {
	return &Theme{Key: key, Light: t.Light.Clone(), Dark: t.Dark.Clone()}
}

func (t *Theme) Equals(x *Theme) bool {
	return t.Light.Equals(x.Light) && t.Dark.Equals(x.Dark)
}

type Themes []*Theme

func (t Themes) Sort() {
	slices.SortFunc(t, func(l *Theme, r *Theme) bool {
		if l.Key == ThemeDefault.Key {
			return true
		}
		if r.Key == ThemeDefault.Key {
			return false
		}
		return l.Key < r.Key
	})
}

func (t Themes) Replace(n *Theme) Themes {
	for idx, o := range t {
		if o.Key == n.Key {
			t[idx] = n
			return t
		}
	}
	ret := append(Themes{}, t...)
	ret = append(ret, n)
	ret.Sort()
	return ret
}

func addLine(sb io.StringWriter, s string, indent int) {
	indention := util.StringRepeat("  ", indent)
	_, _ = sb.WriteString(indention + s + "\n")
}
