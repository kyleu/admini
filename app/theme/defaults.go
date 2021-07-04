package theme

var ThemeDefault = &Theme{
	Key: "default",
	Light: Colors{
		Border: "1px solid #dddddd", LinkDecoration: "none",
		Foreground: "#000000", ForegroundMuted: "#999999",
		Background: "#ffffff", BackgroundMuted: "#eeeeee",
		Link: "#2d414e", LinkVisited: "#406379",
		NavForeground: "#000000", NavBackground: "#4f9abd",
		MenuForeground: "#000000", MenuBackground: "#f0f8ff", MenuBackgroundSelected: "#faebd7",
		ModalBackdrop: "rgba(77, 77, 77, .7)", Success: "#008000", Error: "#FF0000",
	},
	Dark: Colors{
		Border: "1px solid #666666", LinkDecoration: "none",
		Foreground: "#ffffff", ForegroundMuted: "#999999",
		Background: "#121212", BackgroundMuted: "#333333",
		Link: "#dddddd", LinkVisited: "#aaaaaa",
		NavForeground: "#ffffff", NavBackground: "#2d414e",
		MenuForeground: "#dddddd", MenuBackground: "#171f24", MenuBackgroundSelected: "#333333",
		ModalBackdrop: "rgba(33, 33, 33, .7)", Success: "#008000", Error: "#FF0000",
	},
}

var ThemeInverse = &Theme{
	Key: "inverse",
	Light: ThemeDefault.Dark,
	Dark: ThemeDefault.Light,
}
