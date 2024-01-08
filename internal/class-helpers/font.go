package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func FontColor(color string) string {
	className := ClassNameSanitiser(fmt.Sprintf("fc-%s", color))
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"color": color,
		},
	})

	return className
}
