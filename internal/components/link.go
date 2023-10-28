package components

import (
	"strings"

	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

func Link(inner *g.HTMLElement, url string) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	linkClassName := "link"
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

	externalClassName := "external-link"
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + externalClassName + "::after",
		Props: g.CSSProps{
			"content": "\" ↗\"",
		},
	})

	classList := []string{linkClassName}
	target := "_self"

	if strings.HasPrefix(url, "http") {
		classList = append(classList, externalClassName)
		target = "_blank"
	}

	return g.A(g.EB{
		ClassList: classList,
		Children:  g.CE{inner},
		Props: g.Props{
			"href":   url,
			"target": target,
		},
	})
}
