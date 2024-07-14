package theme

import t "github.com/zaptross/portfoligo/internal/types"

var vscl = t.Pallette{
	Base3:   "#1e1e1e",
	Base2:   "#252526",
	Base1:   "#2d2d30",
	Base0:   "#4d4d4d",
	Base00:  "#bbbbbb",
	Base01:  "#cccccc",
	Base02:  "#eeeeee",
	Base03:  "#ffffff",
	Yellow:  "#dda700",
	Orange:  "#E99A66",
	Red:     "#c74e4e",
	Magenta: "#dca3a3",
	Violet:  "#C678DD",
	Blue:    "#006acc",
	Cyan:    "#4EC9B0",
	Green:   "#4dC760",
}

var VSCodeLight = initTheme(
	"vscode-light",
	t.Theme{
		Colors: t.Colors{
			Pallette: vscl,
			Text: t.TextColors{
				Primary:   vscl.Base2,
				Secondary: vscl.Base0,
				Contrast:  vscl.Base03,
				Heading:   vscl.Yellow,
				Link:      vscl.Blue,
				Hover:     vscl.Violet,
			},
			Background: t.BackgroundColors{
				Primary:   vscl.Base03,
				Secondary: vscl.Base02,
				Button:    vscl.Blue,
				Accent:    vscl.Cyan,
				Hover:     vscl.Violet,
			},
			Error:   vscl.Red,
			Warning: vscl.Orange,
			Info:    vscl.Blue,
			Success: vscl.Green,
		},
	},
)
