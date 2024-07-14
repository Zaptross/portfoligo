package types

import (
	"fmt"
)

type Theme struct {
	selector string
	Colors   Colors
	Fonts    Fonts
	Weights  FontWeights
	Borders  Borders
	Radii    Radii
	ZIndices ZIndices
}

type Colors struct {
	Text       TextColors
	Icons      TextColors
	Background BackgroundColors
	Pallette   Pallette
	Filters    Pallette
	Error      string
	Warning    string
	Info       string
	Success    string
}

type TextColors struct {
	Primary   string
	Secondary string
	Contrast  string
	Link      string
	Heading   string
	Hover     string
}

type BackgroundColors struct {
	Primary   string
	Secondary string
	Accent    string
	Button    string
	Hover     string
}

type Fonts struct {
	Body      string
	Heading   string
	Monospace string
}

type FontWeights struct {
	Body    string
	Heading string
	Bold    string
}

type Borders struct {
	None  string
	Thin  string
	Thick string
	Style BorderStyle
}

type BorderStyle struct {
	Default string
}

type Radii struct {
	None    string
	Small   string
	Default string
	Large   string
}

type ZIndices struct {
	Auto    string
	Overlay string
	Drawers string
	Modals  string
}

func (t *Theme) Spacing(i int) string {
	return fmt.Sprintf("%dpx", i*4)
}

func (t *Theme) FontSize(i int) string {
	return fmt.Sprintf("%dpx", i)
}

func (t *Theme) Selector() string {
	return t.selector
}

func (t Theme) Setup(selector string) Theme {
	t.selector = selector
	return t
}

var BaseTheme = Theme{
	Borders: Borders{
		None:  "none",
		Thin:  "1px",
		Thick: "2px",
		Style: BorderStyle{
			Default: "solid",
		},
	},
	Radii: Radii{
		None:    "0",
		Small:   "4px",
		Default: "8px",
		Large:   "16px",
	},
	ZIndices: ZIndices{
		Auto:    "auto",
		Overlay: "100",
		Drawers: "200",
		Modals:  "300",
	},
	Weights: FontWeights{
		Body:    "400",
		Heading: "600",
		Bold:    "700",
	},
}
