package atoms

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

func ImagePlaceholder(width, height int, classList []string) *g.HTMLElement {
	return Image(fmt.Sprintf("https://via.placeholder.com/%dx%d", width, height), classList)
}
