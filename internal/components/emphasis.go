package components

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

func Emphasis(text string) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	className := "emphasis"

	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"font-weight": "bold",
			"color":       theme.Green,
		},
	})

	return g.Em(g.EB{
		ClassList: []string{className},
		Text:      text,
	})
}
