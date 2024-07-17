package atoms

import (
	g "github.com/zaptross/gorgeous"
)

func Image(src string, classList []string) *g.HTMLElement {
	return g.Img(g.EB{
		ClassList: classList,
		Props: g.Props{
			"src": src,
		},
	})
}
