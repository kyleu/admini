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
}

func (t *Theme) CSS(indent int) string {
	sb := &strings.Builder{}
	addLine(sb, "", indent)
	sb.WriteString(t.Light.CSS(indent))
	addLine(sb, "", indent)
	addLine(sb, "@media (prefers-color-scheme: dark) {", indent)
	sb.WriteString(t.Dark.CSS(indent + 1))
	addLine(sb, "}", indent)
	return sb.String()
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
