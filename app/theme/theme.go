package theme

import (
	"io"
	"sort"
	"strings"
)

type Theme struct {
	Key   string `json:"key"`
	Light Colors `json:"light"`
	Dark  Colors `json:"dark"`
	css   string
}

func (t *Theme) CSS() string {
	indent := 0
	if t.css != "" {
		return t.css
	}
	sb := &strings.Builder{}
	sb.WriteString("/* Theme [" + t.Key + "] */\n")
	sb.WriteString(t.Light.CSS(":root", indent))
	sb.WriteString(t.Light.CSS(".mode-light", indent))
	sb.WriteString(t.Dark.CSS(".mode-dark", indent))
	addLine(sb, "", indent)
	addLine(sb, "@media (prefers-color-scheme: dark) {", indent)
	sb.WriteString(t.Dark.CSS(":root", indent + 1))
	sb.WriteString(t.Light.CSS(".mode-light", indent + 1))
	sb.WriteString(t.Dark.CSS(".mode-dark", indent + 1))
	addLine(sb, "}", indent)
	t.css = sb.String()
	return t.css
}

type Themes []*Theme

func (t Themes) Sort() {
	sort.Slice(t, func(i, j int) bool {
		return t[i].Key < t[j].Key
	})
}

func addLine(sb io.StringWriter, s string, indent int) {
	indention := ""
	for i := 0; i < indent; i++ {
		indention += "  "
	}
	_, _ = sb.WriteString(indention + s + "\n")
}
