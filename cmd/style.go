package main

import (
	g "github.com/zaptross/gorgeous"
)

// globalClasses registers some global classes that are used in the project.
func globalClasses() {
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: ".max-wh-100",
		Props: g.CSSProps{
			"max-width":  "100%",
			"max-height": "100%",
		},
	})

	g.Class(&g.CSSClass{
		Selector: ".justify-space-around",
		Props: g.CSSProps{
			"justify-content": "space-around",
		},
	})

	g.Class(&g.CSSClass{
		Selector: ".flex-wrap",
		Props: g.CSSProps{
			"flex-wrap": "wrap",
		},
	})

	g.Class(&g.CSSClass{
		Selector: ".no-flex-shrink",
		Props: g.CSSProps{
			"flex-shrink": "0",
		},
	})
}
