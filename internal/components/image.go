package components

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/types"
)

func ImgProject(p types.PageDetails, image string, classList []string) *g.HTMLElement {
	return ImageRel(fmt.Sprintf("%s/%s", p.GetRelativeURL(), image), classList)
}
func ImageRel(src string, classList []string) *g.HTMLElement {
	return g.Img(g.EB{
		ClassList: classList,
		Props: g.Props{
			"src": fmt.Sprintf("/public%s", src),
		},
	})
}
func Image(src string, classList []string) *g.HTMLElement {
	return g.Img(g.EB{
		ClassList: classList,
		Props: g.Props{
			"src": "https://via.placeholder.com/150",
		},
	})
}
