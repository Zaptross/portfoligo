package main

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

func registerClasses() {
	theme := p.ThemeProvider.GetTheme()

	// Because this class is not used anywhere, it will be tree-shaken from the final css.
	g.Class(&g.CSSClass{
		Selector: ".to-be-tree-shaken",
		Props: g.CSSProps{
			"color": theme.Red,
		},
	})

}
