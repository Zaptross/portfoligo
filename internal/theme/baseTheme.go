package theme

import (
	t "github.com/zaptross/portfoligo/internal/types"
)

var BaseTheme = t.Theme{
	Borders: t.Borders{
		None:  "none",
		Thin:  "1px",
		Thick: "2px",
		Style: t.BorderStyle{
			Default: "solid",
		},
	},
	Fonts: t.Fonts{
		Body:      "'Lato', sans-serif",
		Heading:   "'Roboto', 'Lato', sans-serif",
		Monospace: "'Courier New', monospace",
	},
	Radii: t.Radii{
		None:    "0",
		Small:   "4px",
		Default: "8px",
		Large:   "16px",
	},
	ZIndices: t.ZIndices{
		Auto:    "auto",
		Overlay: "100",
		Drawers: "200",
		Modals:  "300",
	},
	Weights: t.FontWeights{
		Body:    "400",
		Heading: "600",
		Bold:    "700",
	},
}
