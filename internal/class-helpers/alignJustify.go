package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

type Spacing string
type Justifications struct {
	Start, Center, End, SpaceBetween, SpaceAround, SpaceEvenly Spacing
}
type Alignments struct {
	Start, Center, End, Stretch, Baseline, FlexEnd, FlexStart Spacing
}

var Content = Justifications{
	Start:        "start",
	Center:       "center",
	End:          "end",
	SpaceBetween: "space-between",
	SpaceAround:  "space-around",
	SpaceEvenly:  "space-evenly",
}
var Align = Alignments{
	Start:     "start",
	Center:    "center",
	End:       "end",
	Stretch:   "stretch",
	Baseline:  "baseline",
	FlexEnd:   "flex-end",
	FlexStart: "flex-start",
}

func JustifyContent(justification Spacing) string {
	cn := ClassNameSanitiser(fmt.Sprintf("jc-%s", justification))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"justify-content": string(justification),
		},
	})

	return cn
}

func AlignContent(alignment Spacing) string {
	cn := ClassNameSanitiser(fmt.Sprintf("ac-%s", alignment))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"align-content": string(alignment),
		},
	})

	return cn
}
func AlignItems(alignment Spacing) string {
	cn := ClassNameSanitiser(fmt.Sprintf("ac-%s", alignment))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"align-items": string(alignment),
		},
	})

	return cn
}
