package components

import (
	"strings"

	g "github.com/zaptross/gorgeous"
	p "github.com/zaptross/portfoligo/internal/provider"
)

func LinkNav(inner *g.HTMLElement, url string) *g.HTMLElement  { return link(inner, url, "nav", false, false) }
func LinkIcon(inner *g.HTMLElement, url string) *g.HTMLElement { return link(inner, url,  "icon", false, false) }
func Link(inner *g.HTMLElement, url string) *g.HTMLElement     { return link(inner, url, "", true, true) }
func link(inner *g.HTMLElement, url string, differentiator string, showExternal bool, underline bool) *g.HTMLElement {
	theme := p.ThemeProvider.GetTheme()

	linkClassName := "link" + differentiator
	textDeco := "none"
	if underline {
		textDeco = "underline"
	}
	g.Class(&g.CSSClass{
		Selector: "." + linkClassName,
		Props: g.CSSProps{
			"color":           theme.Green,
			"text-decoration": textDeco,
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a:hover",
		Props: g.CSSProps{
			"color":           theme.Violet,
			"text-decoration": "none",
			"transition":      "color 0.25s ease-in-out",
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
		if showExternal {
			classList = append(classList, externalClassName)
		}
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
