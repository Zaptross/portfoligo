package components

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

func P(eb g.EB) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	className := "basic-p"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"color": theme.Base2,
		},
	})

	eb.ClassList = append(eb.ClassList, className)

	return g.P(eb)
}
