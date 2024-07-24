package classhelpers

import (
	"fmt"

	g "github.com/zaptross/gorgeous"
)

const (
	PHONE_ASPECT_BREAKPOINT = "7/9"
)

var mediaPhoneAspect = fmt.Sprintf("(max-aspect-ratio: %s)", PHONE_ASPECT_BREAKPOINT)

func MediaPhone(className string, props g.CSSProps) {
	g.Media(mediaPhoneAspect, "."+className, props)
}
