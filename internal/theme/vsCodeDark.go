package theme

import t "github.com/zaptross/portfoligo/internal/types"

var vscd = t.Pallette{
	Base03:  "#1e1e1e",
	Base02:  "#252526",
	Base01:  "#2d2d30",
	Base00:  "#eaeaea",
	Base0:   "#cccccc",
	Base1:   "#cccccc",
	Base2:   "#ffffff",
	Base3:   "#ffffff",
	Yellow:  "#E5C07B",
	Orange:  "#D19A66",
	Red:     "#c74e4e",
	Magenta: "#dca3a3",
	Violet:  "#C678DD",
	Blue:    "#006acc",
	Cyan:    "#4EC9B0",
	Green:   "#4dC760",
}

var VSCodeDark = initTheme(
	"vscode-dark",
	t.Theme{
		Colors: t.Colors{
			Pallette: vscd,
			Text: t.TextColors{
				Primary:   vscd.Base2,
				Secondary: vscd.Base1,
				Contrast:  vscd.Base03,
				Heading:   vscd.Yellow,
				Link:      vscd.Blue,
				Hover:     vscd.Violet,
			},
			Background: t.BackgroundColors{
				Primary:   vscd.Base03,
				Secondary: vscd.Base02,
				Button:    vscd.Blue,
				Accent:    vscd.Cyan,
				Hover:     vscd.Violet,
			},
			Error:   vscd.Red,
			Warning: vscd.Orange,
			Info:    vscd.Blue,
			Success: vscd.Green,
		},
	},
)
