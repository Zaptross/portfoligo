package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func SimpleIcon(icon string, css g.CSSProps) *g.HTMLElement {
	className := "simple-icon"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"height": "1em",
		},
	})

	return g.Img(g.EB{
		Style:     css,
		ClassList: []string{"filter", className},
		Props: g.Props{
			"src": simpleIconCDNUrl(icon),
		},
	})
}

func simpleIconCDNUrl(icon string) string {
	if icon == "csharp" {
		return "https://cdn.jsdelivr.net/npm/simple-icons@v11/icons/csharp.svg"
	}

	return fmt.Sprintf("https://cdn.jsdelivr.net/npm/simple-icons@v13/icons/%s.svg", icon)
}
