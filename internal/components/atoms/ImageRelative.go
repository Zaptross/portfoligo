package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func ImageRel(src string, classList []string) *g.HTMLElement {
	return g.Img(g.EB{
		ClassList: classList,
		Props: g.Props{
			"src":  fmt.Sprintf("/public%s", src),
			"lazy": "true",
		},
	})
}
