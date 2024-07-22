package components

import (
	g "github.com/zaptross/gorgeous"
	"github.com/zaptross/portfoligo/internal/theme"
)

func Emphasis(text string) *g.HTMLElement {
	t := theme.UseTheme()

	className := "emphasis"

	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"font-weight":  t.Weights.Bold,
			"color":        t.Colors.Text.Link,
			"margin-right": t.Spacing(1),
		},
	})

	return g.Em(g.EB{
		ClassList: []string{className},
		Text:      text,
	})
}
