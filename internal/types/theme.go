package types

import "fmt"

type Theme struct {
	space    int
	Colors   Colors
	Fonts    Fonts
	Weights  FontWeights
	Borders  Borders
	Radii    Radii
	ZIndices ZIndices
}

type Colors struct {
	Text       TextColors
	Background BackgroundColors
	Error      string
	Warning    string
	Info       string
	Success    string
}

type TextColors struct {
	Primary   string
	Secondary string
	Link      string
	Heading   string
	Hover     string
}

type BackgroundColors struct {
	Primary   string
	Secondary string
	Accent    string
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
	return fmt.Sprintf("%dpx", i*t.space)
}

func (t *Theme) FontSize(i int) string {
	return fmt.Sprintf("%dpx", i)
}
