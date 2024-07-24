package atoms

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

const (
	ROOT_STYLE_CLASS = "root-style"
)

func ApplyRootStyle() {
	t := theme.UseTheme()
	g.Class(&g.CSSClass{
		Include:  true,
		Selector: "." + ROOT_STYLE_CLASS,
		Props: g.CSSProps{
			"font-size":   t.FontSize(5),
			"font-family": t.Fonts.Body,
		},
	})

	ch.MediaPhone(ROOT_STYLE_CLASS, g.CSSProps{
		"font-size": t.FontSize(4),
	})
}
