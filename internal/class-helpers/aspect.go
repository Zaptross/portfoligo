package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func Aspect(ratio string) string {
	cn := ClassNameSanitiser(fmt.Sprintf("aspect-%s", ratio))
	g.Class(&g.CSSClass{
		Selector: fmt.Sprintf(".%s", cn),
		Props: g.CSSProps{
			"aspect-ratio": ratio,
		},
	})

	return cn
}
