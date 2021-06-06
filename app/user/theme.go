package user

import (
	"io"
	"strings"
)

type Theme struct {
	Key   string
	Light Colors
	Dark  Colors
}

func (t *Theme) CSS(indent int) string {
	sb := &strings.Builder{}
	add(sb, "", indent)
	sb.WriteString(t.Light.CSS(indent))
	add(sb, "", indent)
	add(sb, "@media (prefers-color-scheme: dark) {", indent)
	sb.WriteString(t.Dark.CSS(indent + 1))
	add(sb, "}", indent)
	return sb.String()
}

var ThemeDefault = &Theme{
	Key: "default",
	Light: Colors{
		Foreground: "#000000", ForegroundMuted: "#999999",
		Background: "#ffffff", BackgroundMuted: "#eeeeee",
		Highlight: "#008000", Link: "#2d414e", LinkVisited: "#406379",
		NavForeground: "#000000", NavBackground: "#4f9abd",
		MenuForeground: "#000000", MenuBackground: "#f0f8ff", MenuBackgroundSelected: "#faebd7",
	},
	Dark: Colors{
		Foreground: "#ffffff", ForegroundMuted: "#999999",
		Background: "#121212", BackgroundMuted: "#333333",
		Highlight: "#008000", Link: "#dddddd", LinkVisited: "#aaaaaa",
		NavForeground: "#ffffff", NavBackground: "#2d414e",
		MenuForeground: "#dddddd", MenuBackground: "#171f24", MenuBackgroundSelected: "#333333",
	},
}

func add(sb io.StringWriter, s string, indent int) {
	indention := ""
	for i := 0; i < indent; i++ {
		indention += "  "
	}
	_, _ = sb.WriteString(indention + s + "\n")
}
