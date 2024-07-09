package types

import "fmt"

type Theme struct {
	space  int
	Colors struct {
		Text struct {
			Primary   string
			Secondary string
			Link      string
			Heading   string
			Hover     string
		}
		Background struct {
			Primary   string
			Secondary string
			Accent    string
		}
		Error   string
		Warning string
		Info    string
		Success string
	}
	Fonts struct {
		Body      string
		Heading   string
		Monospace string
	}
	FontWeights struct {
		Body    string
		Heading string
		Bold    string
	}
	Borders struct {
		None  string
		Thin  string
		Thick string
	}
	Radii struct {
		None    string
		Small   string
		Default string
		Large   string
	}
	ZIndices struct {
		Auto    string
		Overlay string
		Drawers string
		Modals  string
	}
}

func (t *Theme) Spacing(i int) string {
	return fmt.Sprintf("%dpx", i*t.space)
}

func (t *Theme) FontSize(i int) string {
	return fmt.Sprintf("%dpx", i)
}
