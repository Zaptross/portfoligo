package molecules

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/components/atoms"
	"github.com/zaptross/portfoligo/internal/types"
)

func ProjectImage(p types.PageDetails, image string, classList []string) *g.HTMLElement {
	return atoms.ImageRel(fmt.Sprintf("%s/%s", p.GetRelativeURL(), image), classList)
}
