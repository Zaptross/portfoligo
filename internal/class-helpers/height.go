package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func H(height string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("h-%s", height))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"height": height,
		},
	})

	return cn
}

func MxH(height string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("mxh-%s", height))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"max-height": height,
		},
	})

	return cn
}

func MnH(height string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("mnh-%s", height))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"min-height": height,
		},
	})

	return cn
}
