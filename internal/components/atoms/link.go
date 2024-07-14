package atoms

import (
	"strings"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
)

func LinkNav(inner *g.HTMLElement, url string) *g.HTMLElement {
	return link(inner, url, "nav", false, false)
}
func LinkIcon(inner *g.HTMLElement, url string) *g.HTMLElement {
	return link(inner, url, "icon", false, false)
}
func Link(inner *g.HTMLElement, url string) *g.HTMLElement { return link(inner, url, "", true, true) }
func linkHash(inner *g.HTMLElement, url string) *g.HTMLElement {
	return link(inner, url, "hash", true, false)
}
func link(inner *g.HTMLElement, url string, differentiator string, showExternal bool, underline bool) *g.HTMLElement {
	t := theme.UseTheme()

	linkClassName := "link" + differentiator
	textDeco := "none"
	if underline {
		textDeco = "underline"
	}
	g.Class(&g.CSSClass{
		Selector: "." + linkClassName,
		Props: g.CSSProps{
			"color":           t.Colors.Text.Link,
			"text-decoration": textDeco,
		},
	})

	g.Class(&g.CSSClass{
		Selector: "a:hover",
		Props: g.CSSProps{
			"color":           t.Colors.Text.Hover,
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
