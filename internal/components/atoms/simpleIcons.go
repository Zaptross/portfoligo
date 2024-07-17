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
			"src": fmt.Sprintf("https://cdn.jsdelivr.net/npm/simple-icons@v11/icons/%s.svg", icon),
		},
	})
}
