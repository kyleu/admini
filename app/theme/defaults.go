package theme

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
