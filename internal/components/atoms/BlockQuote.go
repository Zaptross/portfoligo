package atoms

import (
	g "github.com/zaptross/gorgeous"
	ch "github.com/zaptross/portfoligo/internal/class-helpers"
	"github.com/zaptross/portfoligo/internal/theme"
)

func BlockQuote(children *g.HTMLElement, color string, classList []string) *g.HTMLElement {
	t := theme.UseTheme()

	className := "block-quote" + ch.ClassNameSanitiser(color)
	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"font-style":  "italic",
			"border-left": "solid thick " + color,
			"margin":      t.Spacing(5),
		},
	})

	return g.Blockquote(g.EB{
		ClassList: append(classList, className, ch.MarginL(t.Spacing(5))),
		Children:  g.CE{children},
	})
}
