package atoms

import (
	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
)

func P(eb g.EB) *g.HTMLElement {
	t := theme.UseTheme()

	className := "basic-p"
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"color": t.Colors.Text.Primary,
		},
	})

	eb.ClassList = append(eb.ClassList, className)

	return g.P(eb)
}
