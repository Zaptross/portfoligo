package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

func Superscript() string {
	selector := "superscript"
	g.Class(&g.CSSClass{
		Selector: "." + selector,
		Props: g.CSSProps{
			"vertical-align": "super",
			"font-size":      "0.6em",
		},
	})

	return selector
}
