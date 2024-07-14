package theme

import t "github.com/zaptross/portfoligo/internal/types"

var sdp = t.Pallette{
	Base03:  "#002b36",
	Base02:  "#073642",
	Base01:  "#586e75",
	Base00:  "#657b83",
	Base0:   "#839496",
	Base1:   "#93a1a1",
	Base2:   "#eee8d5",
	Base3:   "#fdf6e3",
	Yellow:  "#b58900",
	Orange:  "#cb4b16",
	Red:     "#dc322f",
	Magenta: "#d33682",
	Violet:  "#6c71c4",
	Blue:    "#268bd2",
	Cyan:    "#2aa198",
	Green:   "#859900",
}

var SolarizedDark = initTheme(
	"solarized-dark",
	t.Theme{
		Colors: t.Colors{
			Pallette: sdp,
			Text: t.TextColors{
				Primary:   sdp.Base2,
				Secondary: sdp.Base1,
				Contrast:  sdp.Base03,
				Heading:   sdp.Yellow,
				Link:      sdp.Green,
				Hover:     sdp.Violet,
			},
			Background: t.BackgroundColors{
				Primary:   sdp.Base02,
				Secondary: sdp.Base03,
				Button:    sdp.Blue,
				Accent:    sdp.Violet,
				Hover:     sdp.Cyan,
			},
			Error:   sdp.Red,
			Warning: sdp.Orange,
			Info:    sdp.Blue,
			Success: sdp.Green,
		},
	},
)
