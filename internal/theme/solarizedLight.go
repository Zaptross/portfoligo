package theme

import t "github.com/zaptross/portfoligo/internal/types"

var slp = t.Pallette{
	Base03:  "#fdf6e3",
	Base02:  "#eee8d5",
	Base01:  "#93a1a1",
	Base00:  "#839496",
	Base0:   "#657b83",
	Base1:   "#586e75",
	Base2:   "#073642",
	Base3:   "#002b36",
	Yellow:  "#b58900",
	Orange:  "#cb4b16",
	Red:     "#dc322f",
	Magenta: "#d33682",
	Violet:  "#6c71c4",
	Blue:    "#268bd2",
	Cyan:    "#2aa198",
	Green:   "#859900",
}

var SolarizedLight = initTheme(
	"solarized-light",
	t.Theme{
		Colors: t.Colors{
			Pallette: slp,
			Text: t.TextColors{
				Primary:   slp.Base2,
				Secondary: slp.Base1,
				Contrast:  slp.Base03,
				Heading:   slp.Yellow,
				Link:      slp.Green,
				Hover:     slp.Violet,
			},
			Background: t.BackgroundColors{
				Primary:   slp.Base02,
				Secondary: slp.Base03,
				Button:    slp.Blue,
				Accent:    slp.Violet,
				Hover:     slp.Cyan,
			},
			Error:   slp.Red,
			Warning: slp.Orange,
			Info:    slp.Blue,
			Success: slp.Green,
		},
	},
)
