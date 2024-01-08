package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func W(width string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("w-%s", width))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"width": width,
		},
	})

	return cn
}

func MxW(width string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("mxw-%s", width))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"max-width": width,
		},
	})

	return cn
}

func MnW(width string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("mnw-%s", width))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"min-width": width,
		},
	})

	return cn
}
