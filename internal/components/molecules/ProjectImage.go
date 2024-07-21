package molecules

import (
	"fmt"
	"strings"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/types"
)

func ProjectImage(p types.PageDetails, image string, classList []string) *g.HTMLElement {
	if strings.Contains(image, ".mp4") {
		return projectImageMP4(p, image, classList)
	}

	return atoms.ImageRel(fmt.Sprintf("%s/%s", p.GetRelativeURL(), image), classList)
}

func projectImageMP4(p types.PageDetails, image string, classList []string) *g.HTMLElement {
	return g.CustomElement("video", true, g.EB{
		ClassList: classList,
		Props: g.Props{
			"src":      fmt.Sprintf("/public/%s/%s", p.GetRelativeURL(), image),
			"autoplay": "true",
			"loop":     "true",
			"muted":    "true",
			"lazy":     "true",
		},
	})
}
