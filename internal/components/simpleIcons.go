package components

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
			// Use this tool to generate the filter: https://codepen.io/sosuke/pen/Pjoqqp
			"filter": "invert(68%) sepia(13%) saturate(192%) hue-rotate(131deg) brightness(90%) contrast(90%)",
		},
	})
	g.Class(&g.CSSClass{
		Selector: ".linknav:hover ." + className,
		Include:  true,
		Props: g.CSSProps{
			"transition": "filter 0.25s ease-in-out",
			"filter":     "invert(47%) sepia(15%) saturate(1613%) hue-rotate(199deg) brightness(95%) contrast(90%)",
		},
	})

	return g.Img(g.EB{
		Style:     css,
		ClassList: []string{className},
		Props: g.Props{
			"src": fmt.Sprintf("https://cdn.jsdelivr.net/npm/simple-icons@v11/icons/%s.svg", icon),
		},
	})
}
