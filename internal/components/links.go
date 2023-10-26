package components

import (
	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

const (
	linkClassName = "link"
)

func linkClasses() {
	theme := p.ThemeProvider.GetTheme()

	g.Class(&g.CSSClass{
		Selector: "." + linkClassName,
		Props: g.CSSProps{
			"color":           theme.Green,
			"text-decoration": "underline",
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a:hover",
		Props: g.CSSProps{
			"color":           theme.Violet,
			"text-decoration": "none",
		},
	})
}

func LinkRel(text string, relative string) *g.HTMLElement {
	linkClasses()

	return g.A(g.EB{
		ClassList: []string{linkClassName},
		Text:      text,
		Props: g.Props{
			"href": relative,
		},
	})
}

func LinkExternal(text string, url string) *g.HTMLElement {
	linkClasses()

	externalClassName := "external-link"
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + externalClassName + "::after",
		Props: g.CSSProps{
			"content": "\" ↗\"",
		},
	})

	return g.A(g.EB{
		ClassList: []string{linkClassName, externalClassName},
		Text:      text,
		Props: g.Props{
			"href":   url,
			"target": "_blank",
		},
	})
}
