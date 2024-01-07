package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

type Justification string
type Justifications struct {
	Start, Center, End, SpaceBetween, SpaceAround, SpaceEvenly Justification
}

var Content = Justifications{
	Start:        "start",
	Center:       "center",
	End:          "end",
	SpaceBetween: "space-between",
	SpaceAround:  "space-around",
	SpaceEvenly:  "space-evenly",
}

func JustifyContent(justification Justification) string {
	cn := classNameSanitiser(fmt.Sprintf("jc-%s", justification))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"justify-content": string(justification),
		},
	})

	return cn
}
