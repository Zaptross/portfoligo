package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func FilterColor(color string) string {
	className := ClassNameSanitiser(fmt.Sprintf("fil-%s", color))
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"filter": color,
		},
	})

	return className
}
