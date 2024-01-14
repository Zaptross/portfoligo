package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

func Hidden() string {
	className := "hidden"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Include:  true,
		Props: g.CSSProps{
			"display": "none",
		},
	})

	return className
}
